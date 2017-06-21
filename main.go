package main

import (
	"flag"
	"net/http"
	"log"
	"github.com/michaelklishin/rabbit-hole"
)

func main() {
	// get from command ling
	url := flag.String("url", "http://127.0.0.1:15672", "URL of the RabbitMQ broker")
	user := flag.String("user", "guest", "Username of login")
	password := flag.String("password", "guest", "Password of login")
	
	rmqc, err := rabbithole.NewClient(*url, *user, *password)
	qs, err := rmqc.ListQueues()

	if err != nil {
		log.Fatal(err)
	}

	for _, q := range qs {
		resp, err := rmqc.DeleteQueue(q.Vhost, q.Name)

		if err == nil && (resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusNoContent) {
			log.Printf("Queue '%s' in '%s' was deleted!", q.Name, q.Vhost)
		} else {
			if err == nil {
				log.Printf("HTTP Status: %s, error deleting queue '%s' in '%s'", resp.Status, q.Name, q.Vhost)
			} else {
				log.Printf("HTTP OK, but error deleting queue '%s' in '%s': %s", q.Name, q.Vhost, err.Error())
			}
		}
	}

	log.Print("Program completed!")
}
