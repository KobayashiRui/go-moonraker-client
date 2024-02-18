package main

import (
	"context"
	"fmt"

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
		params := moonrakerclient.GetNotifyStatusUpdate(req.Params)
		//fmt.Printf("params: %+v\n", params)
		if val, ok := params["display_status"]; ok {
			//fmt.Printf("exists. The value is %#v", val)
			ds := moonrakerclient.Conv[moonrakerclient.DisplayStatus](val)
			fmt.Printf("Progress: %f\n", ds.Progress*100.0)
		}
		fmt.Println("#################")
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
	init_obj := client.SetSubscribePrinterObject(res)
	fmt.Printf("res: %v\n", init_obj)

	//var idle_timeout_data moonrakerclient.IdleTimeout
	//moonrakerclient.ConvRef(init_obj.Status["idle_timeout"], &idle_timeout_data)

	//fmt.Printf("Get Printer Status: %+v\n", idle_timeout_data)
	//idle_timeout_data2 := moonrakerclient.Conv[moonrakerclient.IdleTimeout](init_obj.Status["idle_timeout"])

	//fmt.Printf("Get Printer Status2: %+v\n", idle_timeout_data2)

	itd := client.GetPrinterObjects([]string{"idle_timeout"})
	fmt.Printf("Get Printer Status3: %+v\n", itd)

	//time.Sleep(time.Second * 3)
	//client.GetServerFilesList()

	//Blocking
	select {}

}
