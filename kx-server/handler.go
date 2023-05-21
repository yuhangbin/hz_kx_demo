package main

import (
	"context"
	management "kx-server/kitex_gen/student/management"
)

// StudentManagementImpl implements the last service interface defined in the IDL.
type StudentManagementImpl struct{}

type StudentInfo struct {
	Num string
	Name string
	Gender string
}

var StudentData = make(map[string]StudentInfo, 5)

// QueryStudent implements the StudentManagementImpl interface.
func (s *StudentManagementImpl) QueryStudent(ctx context.Context, req *management.QueryStudentRequest) (resp *management.QueryStudentResponse, err error) {
	// find student by num
	stu, exist := StudentData[req.Num]
	if !exist {
		return &management.QueryStudentResponse{
			Exist: false,
		}, nil
	}
	return &management.QueryStudentResponse{
		Exist: true,
		Num: stu.Num,
		Name: stu.Name,
		Gender: stu.Gender,
	}, nil
}

// InsertStudent implements the StudentManagementImpl interface.
func (s *StudentManagementImpl) InsertStudent(ctx context.Context, req *management.InsertStudentRequest) (resp *management.InsertStudentResponse, err error) {
	// if this student not in student data then insert into student data
	_, exist := StudentData[req.Num]
	if exist {
		return &management.InsertStudentResponse{
			Ok: false,
			Msg: "This student already exists.",
		}, nil
	}
	StudentData[req.Num] = StudentInfo{
		Num: req.Num,
		Name: req.Name,
		Gender: req.Gender,
	}
	return &management.InsertStudentResponse{
		Ok: true,
	}, nil
}
