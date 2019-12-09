package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	consulapi "github.com/hashicorp/consul/api"
)

type order struct {
	ID    uint64  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Status bool `json:status`
}

type orderConfiguration struct {
	Categories []string `json:"categories"`
}

func registerServiceWithConsul() {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	registration := new(consulapi.AgentServiceRegistration)

	registration.ID = "order"
	registration.Name = "order"
	address := hostname()
	registration.Address = address
	port, err := strconv.Atoi(port()[1:len(port())])
	if err != nil {
		log.Fatalln(err)
	}
	registration.Port = port
	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/healthcheck", address, port)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"
	consul.Agent().ServiceRegister(registration)
}

func Configuration(w http.ResponseWriter, r *http.Request) {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	kvpair, _, err := consul.KV().Get("order-configuration", nil)
	if err != nil {
		fmt.Fprintf(w, "Error. %s", err)
		return
	}
	if kvpair.Value == nil {
		fmt.Fprintf(w, "Configuration empty")
		return
	}
	val := string(kvpair.Value)
	fmt.Fprintf(w, "%s", val)

}

func Orders(w http.ResponseWriter, r *http.Request) {
	orders := []order{
		{
			ID:    1,
			Name:  "Macbook Pro",
			Price: 2000000.00,
			Status: false,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&orders)
}

func main() {
	registerServiceWithConsul()
	http.HandleFunc("/healthcheck", healthcheck)
	http.HandleFunc("/orders", Orders)
	http.HandleFunc("/order/config", Configuration)
	fmt.Printf("order service is up on port: %s", port())
	http.ListenAndServe(port(), nil)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `order service is good`)
}

func port() string {
	p := os.Getenv("order_SERVICE_PORT")
	h := os.Getenv("order_SERVICE_HOST")
	if len(strings.TrimSpace(p)) == 0 {
		return ":8200"
	}
	return fmt.Sprintf("%s:%s", h, p)
}

func hostname() string {
	// return os.Getenv("CONSUL_HTTP_ADDR")
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}
