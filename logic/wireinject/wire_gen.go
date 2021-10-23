// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wireinject

import (
	"gitea.com/liushihao/gostd/internal/data/api/student"
	"gitea.com/liushihao/gostd/internal/data/api/student/class"
	"gitea.com/liushihao/gostd/internal/data/api/student/grades"
	"gitea.com/liushihao/gostd/internal/data/api/student/user-info"
	"gitea.com/liushihao/gostd/internal/data/api/teacher"
	"gitea.com/liushihao/gostd/internal/data/api/teacher/info"
	"gitea.com/liushihao/gostd/internal/data/database"
	"gitea.com/liushihao/gostd/internal/data/database/studentdb"
	"gitea.com/liushihao/gostd/logic"
	"gitea.com/liushihao/gostd/logic/api/handler"
	"gitea.com/liushihao/gostd/logic/api/myrpc"
	"gitea.com/liushihao/gostd/logic/conf"
)

// Injectors from wire.go:

func InitApp(cfg *conf.Cfg) (*logic.App, error) {
	studentDB, err := database.NewStudentDB(cfg)
	if err != nil {
		return nil, err
	}
	table := grades.NewTable(studentDB)
	studentdbStudentDB, err := studentdb.NewStudentDB(cfg)
	if err != nil {
		return nil, err
	}
	userInfoRepo := studentdb.NewUserInfoRepo(studentdbStudentDB)
	classRepo := studentdb.NewClassRepo(studentdbStudentDB)
	dao := studentdb.NewStudentDao(cfg, studentdbStudentDB, userInfoRepo, classRepo)
	cli := userinfo.NewCli(dao)
	classTable := class.NewTable(studentDB)
	api := student.NewApi(table, cli, classTable)
	teacherDB, err := database.NewTeacherDB(cfg)
	if err != nil {
		return nil, err
	}
	infoAPI := info.NewTable(teacherDB)
	teacherAPI := teacher.NewApi(infoAPI)
	server := handler.NewServer(api, teacherAPI)
	myrpcServer := myrpc.NewServer(cfg)
	app := logic.NewApp(cfg, server, api, teacherAPI, myrpcServer)
	return app, nil
}
