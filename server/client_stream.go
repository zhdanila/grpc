package main

import (
	"grpc/proto"
	"io"
	"log"
)

func(s *helloServer) SayHelloClientStreaming(stream proto.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(
				&proto.MessagesList{
					Messages: messages,
				},
			)
		}

		if err != nil {
			return err
		}
		log.Printf("got request with name: %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}