package main

import (
	"bufio"
	"golang.org/x/net/context"
	"log"
	"google.golang.org/grpc"
	pb "../server/src/simple"
	"os"
)

func main(){
	conn,err:=grpc.Dial("localhost:12345", grpc.WithInsecure())
	if err!= nil{
		log.Fatal(err)
	}

	c:=pb.NewSimpleClient(conn)
	in:=bufio.NewReader(os.Stdin)
	for {
		line,_,err:=in.ReadLine()
		if err!= nil{
			log.Fatal(err)
		}

		reply,err:=c.GetLine(context.Background(),&pb.SimpleRequest{Data: string(line)})
		if err!=nil{
			log.Fatal(err)
		}

		log.Printf("Reply: %v, Data: %v",reply,reply.Data)
	}
}
