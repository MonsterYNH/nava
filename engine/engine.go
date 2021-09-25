package engine

import (
	"fmt"
	"net/http"

	"github.com/MonsterYNH/nava/datasource"
	"github.com/MonsterYNH/nava/logger"
	"github.com/MonsterYNH/nava/setting"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Engine struct {
	setting.Config
	datasource.DataSource
	logrus.Logger
	gin      *gin.Engine
	apiGroup *gin.RouterGroup
	services map[string]Service
}

func New(config *setting.Config) (*Engine, error) {
	dataSource, err := datasource.InitDataSource(config)
	if err != nil {
		return nil, err
	}

	logger, err := logger.InitLogger(&config.Logger)
	if err != nil {
		return nil, err
	}

	engine, err := initGin(config, logger)
	if err != nil {
		return nil, err
	}

	return &Engine{
		Config:     *config,
		DataSource: dataSource,
		Logger:     *logger,
		gin:        engine,
		apiGroup:   engine.Group("/api", AuthMiddleware(*config), RecoverMiddleware(*config, logger)),
		services:   make(map[string]Service),
	}, nil
}

func (engine *Engine) RegisterService(prefix string, service Service) error {
	if _, exist := engine.services[service.Name()]; exist {
		return fmt.Errorf("service %s is already exist", service.Name())
	}

	group := engine.gin.Group(fmt.Sprintf("/%s/%s", prefix, service.Name()), service.Middlewares()...)

	return service.RegisterHandler(engine, group)
}

func (engine *Engine) RegisterAPIService(prefix string, service Service) error {
	if _, exist := engine.services[service.Name()]; exist {
		return fmt.Errorf("service %s is already exist", service.Name())
	}

	return service.RegisterHandler(engine, engine.apiGroup.Group(prefix, service.Middlewares()...))
}

func (engine *Engine) GetConfig() setting.Config {
	return engine.Config
}

func Run(engine *Engine) error {
	routesInfo := engine.gin.Routes()
	for _, routeInfo := range routesInfo {
		engine.Infof("method: %s, path: %s, handler: %s\n", routeInfo.Method, routeInfo.Path, routeInfo.Handler)
	}

	engine.Infof(fmt.Sprintf("[INFO] server start on %s:%d", engine.Config.Server.Host, engine.Server.Port))

	return engine.gin.Run(fmt.Sprintf("%s:%d", engine.Config.Server.Host, engine.Config.Server.Port))
}

func initGin(config *setting.Config, logger *logrus.Logger) (*gin.Engine, error) {
	// gin.SetMode should call before gin.New
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()

	if config.Server.EnablePPROF {
		pprof.Register(app)
	}

	if config.Server.EnableHealthCheck {
		app.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, APIResponse{
				ErrCode: ErrSuccess,
			})
		})
	}

	if config.Server.EnableSwagger {
		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return app, nil
}
