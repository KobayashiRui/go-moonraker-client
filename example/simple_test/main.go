package main

import (
	"context"
	"fmt"
	"time"

	moonrakerclient "github.com/KobayashiRui/go-moonraker-client"

	"github.com/sourcegraph/jsonrpc2"
)

func OriginHandler(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {

	switch req.Method {
	//case "notify_proc_stat_update":
	//	var notify_proc_stat_update []moonrakerclient.NotifyProcStatUpdate
	//	err := json.Unmarshal(*req.Params, &notify_proc_stat_update)
	//	if err != nil {
	//		fmt.Println("ERROR notify_proc_stat_update")
	//		fmt.Printf("%v\n", err)
	//	}
	//	fmt.Printf("%v\n", notify_proc_stat_update)
	case "notify_status_update":
		fmt.Println("Notify Status")
		fmt.Printf("%+v\n", string(*req.Params))
		//default:
		//	fmt.Println("other sub")
		//	fmt.Printf("%v\n", string(*req.Params))
	}
}

func main() {
	client := moonrakerclient.NewMoonrakerClient(
		"ws://localhost:7125/websocket")

	client.SetHandler(OriginHandler)
	client.Connect()

	//client.SetSubscribePrinterObject()

	//time.Sleep(time.Second * 3)
	res := client.GetPrinterObjectsList()
	client.SetSubscribePrinterObject(res)

	time.Sleep(time.Second * 3)
	client.GetServerFilesList()

	//Blocking
	select {}

}
