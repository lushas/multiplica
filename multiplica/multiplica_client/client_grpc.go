/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"log"
	"os"
	"time"
	"strconv"
	"fmt"
	"google.golang.org/grpc"
	pb "multiplica/multiplica"
)

const (
	address     = "localhost:9500"
)

func main() {
	// Set up a connection to the server.
	var number1, number2 int
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	m := pb.NewMultiplicarClient(conn)

	// Contact the server and print out its response.
	if len(os.Args) == 3 {
		number1,_ = strconv.Atoi(os.Args[1])
		number2,_ = strconv.Atoi(os.Args[2])
	}else {
		fmt.Println("Wrong parameter: ",len(os.Args))
	}
	n1 := int32(number1)
	n2 := int32(number2)
	fmt.Printf("N1: %d, N2 %d", n1, n2)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := m.Multiplica(ctx, &pb.MessageIn{Number1: n1, Number2: n2})
	if err != nil {
		log.Fatalf("could not make product: %v", err)
	}
	log.Printf("Result: %d", r.Result)
}
