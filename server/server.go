package main

import (
	"context"
	"fmt"
	"net"

	"github.com/bkpeh/flatbuffer/fbs/messages"
	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc"
)

type fserver struct{}

func (s fserver) GetPerson(c context.Context, n *messages.Name) (*flatbuffers.Builder, error) {
	fclient := flatbuffers.NewBuilder(1024)

	fname := fclient.CreateString("Alex")
	lname := fclient.CreateString("Tan")

	messages.NameStart(fclient)
	messages.NameAddLast(fclient, lname)
	messages.NameAddFirst(fclient, fname)
	pname := messages.NameEnd(fclient)

	address := fclient.CreateString("Singapore")
	messages.AddressStart(fclient)
	messages.AddressAddAdd(fclient, address)
	messages.AddressAddPhone(fclient, 1234567)
	add := messages.AddressEnd(fclient)

	messages.PersonStartAgeVector(fclient, 1)
	messages.CreateDOB(fclient, 1990, 5, 1)
	ag := fclient.EndVector(1)

	messages.PersonStart(fclient)
	messages.PersonAddFname(fclient, pname)
	messages.PersonAddAge(fclient, ag)
	messages.PersonAddId(fclient, 30)
	messages.PersonAddDp(fclient, messages.DeptD1)
	messages.PersonAddLocation(fclient, add)
	off := messages.PersonEnd(fclient)
	fclient.Finish(off)

	fmt.Println("GetPerson")

	return fclient, nil
}

func main() {

	lsvr, err := net.Listen("tcp", ":9005")
	if err != nil {
		fmt.Println("Fail to start listening.", err)
	}

	var svr fserver

	gsvr := grpc.NewServer(grpc.CustomCodec(flatbuffers.FlatbuffersCodec{}))

	messages.RegisterGetInfoServer(gsvr, svr)

	if err := gsvr.Serve(lsvr); err != nil {
		fmt.Println("Fail to start grpc.", err)
	}
}
