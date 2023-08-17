package prototest

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func Errorer() {
	e := MyError{
		Status: 1,
		Reason: "Something went wrong",
	}

	anyMessage, err := anypb.New(e)
	if err != nil {
		panic(err)
	}

	m := &SampleMessage{
		Error: anyMessage,
	}

	wire, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	r := new(SampleMessage)
	if err := proto.Unmarshal(wire, r); err != nil {
		panic(err)
	}

	fmt.Printf("Type of Error: %T\n", r.Error)

	er, err := r.Error.UnmarshalNew()
	if err != nil {
		panic(err)
	}

	var errorReceived error
	switch ea := er.(type) {
	case *prototest.MyError:
		errorReceived = ea
	}

	fmt.Printf("type of error after unmarshaling any: %T", errorReceived)
}
