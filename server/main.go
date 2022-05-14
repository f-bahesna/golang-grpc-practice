package main

import (
	"context"
	"encoding/json"
	pb "fbahesna.com/learn/grpc-practice/student"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net"
	"sync"
)

//karena interface, struct harus implementasikan interface dari student.proto disini
type dataStudentServer struct {
	pb.UnimplementedDataStudentServer
	mu       sync.Mutex
	students []*pb.Student
}

func (d *dataStudentServer) FindStudentByEmail(ctx context.Context, student *pb.Student) (*pb.Student, error) {
	for _, v := range d.students {
		if v.Email == student.Email {
			return v, nil
		}
	}

	return nil, nil
}

//nge loa d data students
func (d *dataStudentServer) loadData() {
	data, err := ioutil.ReadFile("data/students.json")
	if err != nil {
		log.Fatalln("error in read file", err.Error())
	}

	fmt.Println(json.Unmarshal(data, &d.students))
	//unmarshal data taruh ke pointer &d.students
	if err := json.Unmarshal(data, &d.students); err != nil {
		log.Fatalln("error in unmarshal gan", err.Error())
	}
}

func newServer() *dataStudentServer {
	s := dataStudentServer{}
	s.loadData()
	return &s
}

func main() {
	//gawe grpc server e ri!
	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalln("Error listening server", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())

	//dilakokne server e ri!
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("error when serve", err.Error())
	}

	fmt.Println("server started tcp:1200")
}
