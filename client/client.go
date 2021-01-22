package main

import (
	"context"
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"

	"github.com/bkpeh/flatbuffer/fbs/messages"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial("192.168.56.1:9005", grpc.WithInsecure(), grpc.WithCodec(flatbuffers.FlatbuffersCodec{}))

	if err != nil {
		fmt.Println("Fail to dial.", err)
	}

	defer conn.Close()

	fclient := flatbuffers.NewBuilder(1024)

	first := fclient.CreateString("Alex")
	last := fclient.CreateString("Tan")

	messages.NameStart(fclient)
	messages.NameAddLast(fclient, last)
	messages.NameAddFirst(fclient, first)
	//off := messages.NameEnd(fclient)
	fclient.Finish(messages.NameEnd(fclient))

	//fmt.Println("Offset:", off)

	client := messages.NewGetInfoClient(conn)

	respond, err := client.GetPerson(context.Background(), fclient)

	if err != nil {
		fmt.Println("Error from respond.", err)
	}

	var add messages.Address
	var d messages.DOB
	respond.Age(&d, 0)

	fmt.Println("Respond:", string(respond.Fname(&messages.Name{}).First()))
	fmt.Println("Respond:", string(respond.Fname(&messages.Name{}).Last()))
	fmt.Println("Respond:", d.Year())
	fmt.Println("Respond:", d.Month())
	fmt.Println("Respond:", d.Day())
	fmt.Println("Respond:", respond.Id())
	fmt.Println("Respond:", respond.Dp())
	fmt.Println("Respond:", string(respond.Location(&add).Add()))
	fmt.Println("Respond:", respond.Location(&add).Phone())
}
