package main

import (
	moonrakerclient "go-moonraker-client"
	"time"
)

func main() {
	client := moonrakerclient.NewMoonrakerClient("ws://localhost:7125/websocket")
	client.Connect()

	time.Sleep(time.Second * 3)
	client.GetPrinterObjectsList()

	time.Sleep(time.Second * 3)
	client.GetServerFilesList()

	//Blocking
	select {}

}
