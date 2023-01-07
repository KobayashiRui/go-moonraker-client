package moonrakerclient

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
	websocketjsonrpc2 "github.com/sourcegraph/jsonrpc2/websocket"
)

type MoonrakerClient struct {
	uri  string
	jrpc *jsonrpc2.Conn
}

func NewMoonrakerClient(uri string) *MoonrakerClient {
	return &MoonrakerClient{uri: uri}
}

func (mc *MoonrakerClient) Test() {
	fmt.Println("Test")
}

func (mc *MoonrakerClient) Connect() {
	ctx := context.Background()
	fmt.Printf("signal uri:%v\n", mc.uri)

	//TODO Error対処
	c, _, err := websocket.DefaultDialer.Dial(mc.uri, nil)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	stream := websocketjsonrpc2.NewObjectStream(c)
	mc.jrpc = jsonrpc2.NewConn(ctx, stream, mc)
}

func (mc *MoonrakerClient) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	fmt.Printf("method: %v \n", req.Method)
	switch req.Method {
	case "notify_proc_stat_update":
		var notify_proc_stat_update []NotifyProcStatUpdate
		err := json.Unmarshal(*req.Params, &notify_proc_stat_update)
		if err != nil {
			fmt.Println("ERROR notify_proc_stat_update")
			fmt.Printf("%v\n", err)
		}
		fmt.Printf("%v\n", notify_proc_stat_update)
	default:
		fmt.Printf("%v\n", string(*req.Params))
	}

}

func (mc *MoonrakerClient) GetPrinterObjectsList() {
	ctx := context.Background()
	var po PrinterObjects

	var params struct{}
	if err := mc.jrpc.Call(ctx, "printer.objects.list", params, &po); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get")
	fmt.Printf("%v \n", po)
}

func (mc *MoonrakerClient) GetServerFilesList() {
	ctx := context.Background()
	var gcode_files []GcodeFileData
	var params struct{}

	if err := mc.jrpc.Call(ctx, "server.files.list", params, &gcode_files); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get Server Files List")
	fmt.Printf("%v \n", gcode_files)
}
