//go:build wireinject
// +build wireinject

package wire

import (
	"cheemshappy_pay/internal/handler"
	"cheemshappy_pay/internal/repository"
	"cheemshappy_pay/internal/server"
	"cheemshappy_pay/internal/service"
	"cheemshappy_pay/internal/task"
	"cheemshappy_pay/pkg/app"
	"cheemshappy_pay/pkg/chain"
	"cheemshappy_pay/pkg/helper/sid"
	"cheemshappy_pay/pkg/jwt"
	"cheemshappy_pay/pkg/log"
	"cheemshappy_pay/pkg/server/http"

	"github.com/google/wire"
	"github.com/imroc/req/v3"
	"github.com/spf13/viper"
)

var repositorySet = wire.NewSet(
	repository.NewDB,
	repository.NewRedis,
	repository.NewRepository,
	repository.NewTransaction,
	repository.NewUserRepository,
	repository.NewWalletRepository,
	repository.NewMerchantsRepository,
	repository.NewMerchantsMetaRepository,
	repository.NewOrderRepository,
	repository.NewMerchantsApiRepository,
	repository.NewSysWalletRepository,
	repository.NewSysConfigRepository,
)

var serviceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
	service.NewWalletService,
	service.NewMerchantsApiService,
	service.NewMerchantsMetaService,
	service.NewMerchantsService,
	service.NewOrderService,
	service.NewSysWalletService,
	service.NewStatsService,
	service.NewSysConfigService,
)

var handlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
	handler.NewWalletHandler,
	handler.NewMerchantsHandler,
	handler.NewMerchantsMetaHandler,
	handler.NewOrderHandler,
	handler.NewMerchantsApiHandler,
	handler.NewSysWalletHandler,
	handler.NewStatsHandler,
	handler.NewSysConfigHandler,
)

var serverSet = wire.NewSet(
	server.NewHTTPServer,
	server.NewJob,
	server.NewTaskServer,
)
var taskSet = wire.NewSet(
	task.NewTask,
	task.NewOrderTask,
)

// build App
func newApp(httpServer *http.Server, job *server.Job, task *server.TaskServer) *app.App {
	return app.NewApp(
		app.WithServer(httpServer, job, task),
		app.WithName("demo-server"),
	)
}

func NewWire(*viper.Viper, *log.Logger) (*app.App, func(), error) {

	panic(wire.Build(
		repositorySet,
		serviceSet,
		handlerSet,
		serverSet,
		sid.NewSid,
		jwt.NewJwt,
		newApp,
		req.NewClient,
		chain.NewVerifierFactory,
		taskSet,
	))
}
