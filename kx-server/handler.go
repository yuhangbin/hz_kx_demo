package main

import (
	"context"
	management "kx-server/kitex_gen/student/management"
)

// StudentManagementImpl implements the last service interface defined in the IDL.
type StudentManagementImpl struct{}

// QueryStudent implements the StudentManagementImpl interface.
func (s *StudentManagementImpl) QueryStudent(ctx context.Context, req *management.QueryStudentRequest) (resp *management.QueryStudentResponse, err error) {
	// TODO: Your code here...
	return
}

// InsertStudent implements the StudentManagementImpl interface.
func (s *StudentManagementImpl) InsertStudent(ctx context.Context, req *management.InsertStudentRequest) (resp *management.InsertStudentResponse, err error) {
	// TODO: Your code here...
	return
}
