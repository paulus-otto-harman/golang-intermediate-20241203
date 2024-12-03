package infra

import (
	"go.uber.org/zap"
	"project/class/config"
	"project/class/database"
	"project/class/handler"
	"project/class/log"
	"project/class/middleware"
	"project/class/repository"
	"project/class/service"
)

type ServiceContext struct {
	Cacher     database.Cacher
	Cfg        config.Config
	Ctl        handler.Handler
	Log        *zap.Logger
	Middleware middleware.Middleware
}

func NewServiceContext(migrateDb bool, seedDb bool) (*ServiceContext, error) {

	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	config, err := config.LoadConfig(migrateDb, seedDb)
	if err != nil {
		handlerError(err)
	}

	// instance looger
	log, err := log.InitZapLogger(config)
	if err != nil {
		handlerError(err)
	}

	// instance database
	db, err := database.ConnectDB(config)
	if err != nil {
		handlerError(err)
	}

	rdb := database.NewCacher(config, 60*60)

	// instance repository
	repository := repository.NewRepository(db)

	// instance service
	service := service.NewService(repository)

	// instance controller
	Ctl := handler.NewHandler(service, log, rdb)

	mw := middleware.NewMiddleware(rdb)

	return &ServiceContext{Cacher: rdb, Cfg: config, Ctl: *Ctl, Log: log, Middleware: mw}, nil
}
