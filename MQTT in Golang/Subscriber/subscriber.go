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
	connection.SetClientID("mqtt-go-subscriber")
	connection.SetDefaultPublishHandler(f)
	subscriber := MQTT.NewClient(connection)
	if token := subscriber.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := subscriber.Subscribe("mqtt-go", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	time.Sleep(10 * time.Second)

	if token := subscriber.Unsubscribe("mqtt-go"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}

	subscriber.Disconnect(250)
}
