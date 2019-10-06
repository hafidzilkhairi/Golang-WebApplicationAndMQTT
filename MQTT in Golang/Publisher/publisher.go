package main

import (
	"fmt"
	"time"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Topic: %s\n", msg.Topic())
	fmt.Printf("Message: %s\n", msg.Payload())
}

func main() {
	connection := MQTT.NewClientOptions().AddBroker("localhost:1883")
	connection.SetClientID("mqtt-go-publisher")
	connection.SetDefaultPublishHandler(f)
	publisher := MQTT.NewClient(connection)
	if token := publisher.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for(true){
		text := fmt.Sprintf("I have reached a point that i can make youtube from this app")
		token := publisher.Publish("mqtt-go", 0, false, text)
		token.Wait()
		time.Sleep( 1 * time.Second)
	}
	publisher.Disconnect(250)
}