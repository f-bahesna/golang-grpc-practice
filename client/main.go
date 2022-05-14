package main

import (
	"context"
	pb "fbahesna.com/learn/grpc-practice/student"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func findStudentDataByEmail(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// main function protobuf
	//pemanggilan hanya menggunakan method, tidak seperti rest http. harus setup header dll.
	s := pb.Student{Email: email}
	student, err := client.FindStudentByEmail(ctx, &s)
	if err != nil {
		log.Fatalln("error find student by email", err.Error())
	}

	fmt.Println(student)
}

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":1200", opts...)
	if err != nil {
		log.Fatalln("error in dial")
	}

	defer conn.Close()

	client := pb.NewDataStudentClient(conn)
	findStudentDataByEmail(client, "frada@bahesna.com")
	findStudentDataByEmail(client, "bahesna@fbahesna.com")
}
