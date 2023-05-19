package main

import (
	management "kx-server/kitex_gen/student/management/studentmanagement"
	"log"
)

func main() {
	svr := management.NewServer(new(StudentManagementImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
