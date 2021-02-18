package main

import (
	"fmt"
	"interview/internal/configs"
	"interview/pkg/client"
	"interview/pkg/database"
	"log"
	"os"
	"os/signal"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func main() {
	//file with DSN and topic
	config := configs.NewConfig()
	//Struct with sql and mqtt client
	client := client.NewClient()

	//struct with sqlite db
	sqlite, err := database.MakeSQlite(config.DNS)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlite.DB.Close()
	client.MyDB = sqlite

	client.SubscribeToTopic(config.Topic)

	//gracefull shutdown
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt)

	//signal to stop listening
	log.Println("Start working")
	<-interruption
	log.Println("\nInterruption")

	client.MQTTClient.Disconnect(250)

}
