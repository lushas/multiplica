package main_test

import(
	"context"
	"testing"
	pb "multiplica/multiplica"
	"google.golang.org/grpc"
	"log"
	"time"

)
const (
        address     = "localhost:9500"
)

func TestMultiplica(t *testing.T) {
	type test struct {
		name string
		req *pb.MessageIn
		result int32
		expectedErr bool
	}
	testCase := new(test)
	testCase.name = "Multiplica server test"
	testCase.req = &pb.MessageIn{Number1: 3, Number2: 5}
	testCase.result = 15
	testCase.expectedErr= false

	t.Run(testCase.name, func(t *testing.T) {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
         	       log.Fatalf("did not connect: %v", err)
        	}
		defer conn.Close()
	        m := pb.NewMultiplicarClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        	defer cancel()
		response, err := m.Multiplica(ctx, testCase.req)
		if ! testCase.expectedErr {
			if response.Result != testCase.result {
				t.Errorf("Multiplica result = %d, want %d", response.Result, testCase.result)
			}
		}
	})
}
