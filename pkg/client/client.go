package client

import (
	"encoding/json"
	"fmt"
	"interview/pkg/database"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	topicTestHandler func(client mqtt.Client, msg mqtt.Message)
)

//Client keeps mqtt client and db for writing
type Client struct {
	MQTTClient mqtt.Client
	MyDB       *database.SQLite
}

//NewClient returns client mqtt client set
func NewClient() *Client {
	c := new(Client)

	broker := "localhost"
	port := 1883

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("emqx")
	opts.SetPassword("public")
	opts.SetDefaultPublishHandler(c.messageHandler)
	opts.OnConnect = c.connectHandler
	opts.OnConnectionLost = c.connectionLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	topicTestHandler = c.topicTestHandler()
	c.MQTTClient = client

	return c
}

func (c *Client) messageHandler(client mqtt.Client, msg mqtt.Message) {
	log.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	switch msg.Topic() {
	case "test/topic":
		topicTestHandler(client, msg)
	}
}

func (c *Client) connectHandler(client mqtt.Client) {
	log.Println("Connected")
}

func (c *Client) connectionLostHandler(client mqtt.Client, err error) {
	log.Printf("Connect lost: %v", err)
}

func (c *Client) topicTestHandler() func(mqtt.Client, mqtt.Message) {
	type test struct {
		Message string `json:"message"`
		Name    string `json:"name"`
	}
	return func(client mqtt.Client, msg mqtt.Message) {
		t := new(test)
		err := json.Unmarshal(msg.Payload(), &t)
		if err != nil {
			log.Printf("Could not proccess %s\n", msg.Payload())
			return
		}
		err = c.MyDB.TestHandler(t.Message, t.Name)
		if err != nil {
			log.Printf("Could not write %s\n", msg.Payload())
			return
		}
	}
}

//SubscribeToTopic subscribes to topic with timeout
func (c *Client) SubscribeToTopic(topicToSubscribe string) {
	topic := topicToSubscribe
	token := c.MQTTClient.Subscribe(topic, 1, nil)
	token.Wait()
	log.Printf("Subscribed to topic: %s", topic)
}
