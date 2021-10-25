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
	"gitea.com/liushihao/gostd/internal/data/database/studentdb"
	"gitea.com/liushihao/gostd/internal/data/database/teacherdb"
	"gitea.com/liushihao/gostd/logic"
	"gitea.com/liushihao/gostd/logic/api/handler"
	"gitea.com/liushihao/gostd/logic/api/myrpc"
	"gitea.com/liushihao/gostd/logic/conf"
)

// Injectors from wire.go:

func InitApp(cfg *conf.Cfg) (*logic.App, error) {
	studentDB, err := studentdb.NewStudentDB(cfg)
	if err != nil {
		return nil, err
	}
	userInfoRepo := studentdb.NewUserInfoRepo(studentDB)
	classRepo := studentdb.NewClassRepo(studentDB)
	dao := studentdb.NewDao(cfg, studentDB, userInfoRepo, classRepo)
	cli := grades.NewTable(dao)
	userinfoCli := userinfo.NewCli(dao)
	classCli := class.NewCli(dao)
	api := student.NewApi(cli, userinfoCli, classCli)
	teacherDB, err := teacherdb.NewTeacherDB(cfg)
	if err != nil {
		return nil, err
	}
	infoRepo := teacherdb.NewInfoRepo(teacherDB)
	teacherdbDao := teacherdb.NewDao(cfg, teacherDB, infoRepo)
	infoCli := info.NewCli(teacherdbDao)
	teacherAPI := teacher.NewApi(infoCli)
	server := handler.NewServer(api, teacherAPI)
	myrpcServer := myrpc.NewServer(cfg)
	app := logic.NewApp(cfg, server, api, teacherAPI, myrpcServer)
	return app, nil
}
