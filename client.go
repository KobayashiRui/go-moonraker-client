package moonrakerclient

import (
	"context"
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
	websocketjsonrpc2 "github.com/sourcegraph/jsonrpc2/websocket"
)

type MoonrakerClient struct {
	uri     string
	jrpc    *jsonrpc2.Conn
	handler func(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request)
}

func NewMoonrakerClient(uri string) *MoonrakerClient {
	return &MoonrakerClient{
		uri:  uri,
		jrpc: nil,
	}
}

func (mc *MoonrakerClient) SetHandler(handler func(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request)) {
	mc.handler = handler
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
	mc.handler(ctx, conn, req)
}

func (mc *MoonrakerClient) GetPrinterObjectsList() PrinterObjects {
	ctx := context.Background()
	var po PrinterObjects

	var params struct{}
	if err := mc.jrpc.Call(ctx, "printer.objects.list", params, &po); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Get")
	fmt.Printf("%v \n", po)

	return po
}

func (mc *MoonrakerClient) GetPrinterObjects(objects []string) PrinterObjectsQuery {
	ctx := context.Background()
	var params map[string]interface{} = make(map[string]interface{})
	var objects_param map[string]*string = make(map[string]*string)
	for _, v := range objects {
		objects_param[v] = nil
	}

	params["objects"] = objects_param

	var res PrinterObjectsQuery
	if err := mc.jrpc.Call(ctx, "printer.objects.query", params, &res); err != nil {
		fmt.Println(err)
	}

	return res
}

func (mc *MoonrakerClient) SetSubscribePrinterObject(po PrinterObjects) PrinterObjectsQuery {
	ctx := context.Background()
	var params map[string]interface{} = make(map[string]interface{})
	var objects_param map[string]*string = make(map[string]*string)
	for _, v := range po.Objects {
		objects_param[v] = nil
	}

	params["objects"] = objects_param

	var res PrinterObjectsQuery
	if err := mc.jrpc.Call(ctx, "printer.objects.subscribe", params, &res); err != nil {
		fmt.Println(err)
	}

	return res
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

func (mc *MoonrakerClient) SendCall(ctx context.Context, method string, params interface{}) (interface{}, error) {

	var res interface{}
	if err := mc.jrpc.Call(ctx, method, params, &res); err != nil {
		fmt.Println(err)
		return res, err
	}
	return res, nil
}
