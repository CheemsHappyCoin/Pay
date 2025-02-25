package server

import (
	"cheemshappy_pay/internal/handler"
	"cheemshappy_pay/internal/middleware"
	"cheemshappy_pay/pkg/jwt"
	"cheemshappy_pay/pkg/log"
	"cheemshappy_pay/pkg/server/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
	walletHandler *handler.WalletHandler,
	merchantsHandler *handler.MerchantsHandler,
	merchantsMetaHandler *handler.MerchantsMetaHandler,
	orderHandler *handler.OrderHandler,
	sysWalletHandler *handler.SysWalletHandler,
	statsHandler *handler.StatsHandler,
	merchantsApiHandler *handler.MerchantsApiHandler,
	sysConfigHandler *handler.SysConfigHandler,
) *http.Server {
	// gin.SetMode(gin.ReleaseMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/login", userHandler.Login)
		}
		// Non-strict permission routing group
		noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
		{
			noStrictAuthRouter.GET("/isLogin", userHandler.IsLogin)
		}

		// Strict permission routing group
		dashboardAuthRouter := v1.Group("/dashboard").Use(middleware.StrictAuth(jwt, logger), middleware.AdminOnly())
		{
			dashboardAuthRouter.GET("/stats", statsHandler.GetDashboardStats)
		}

		strictAuthRouter := v1.Group("/user").Use(middleware.StrictAuth(jwt, logger), middleware.AdminOnly())
		{
			strictAuthRouter.GET("/profile", userHandler.GetProfile)
			strictAuthRouter.POST("/update", userHandler.UpdateUserInfo)
		}

		merAuthRouter := v1.Group("/merchants").Use(middleware.StrictAuth(jwt, logger), middleware.AdminOnly())
		{
			merAuthRouter.GET("/:id", merchantsHandler.GetMerchants)
			merAuthRouter.POST("/list", merchantsHandler.GetMerchantsList)
			merAuthRouter.POST("/create", merchantsHandler.CreateMerchants)
			merAuthRouter.DELETE("/delete/:id", merchantsHandler.DeleteMerchants)
			merAuthRouter.POST("/update", merchantsHandler.UpdateMerchants)
			merAuthRouter.GET("/in", merchantsHandler.GetMerchantsIN)
		}

		merMetaAuthRouter := v1.Group("/merchants/meta").Use(middleware.StrictAuth(jwt, logger), middleware.AdminOnly())
		{
			merMetaAuthRouter.GET("/:mid", merchantsMetaHandler.GetMerchantsMeta)
			merMetaAuthRouter.POST("/update/:mid", merchantsMetaHandler.UpdateMerchantsMeta)
		}
		merWalletAuthRouter := v1.Group("/wallet").Use(middleware.StrictAuth(jwt, logger), middleware.AdminOnly())
		{
			merWalletAuthRouter.GET("/:mid", walletHandler.GetWalletByMID)
			merWalletAuthRouter.POST("/create", sysWalletHandler.DeriveChildWallet)

			merWalletAuthRouter.POST("/add", walletHandler.AddWallet)
			merWalletAuthRouter.DELETE("/delete/:id", walletHandler.DeleteWallet)
			merWalletAuthRouter.POST("/update", walletHandler.UpdateWallet)
			merWalletAuthRouter.POST("/list", walletHandler.GetWallets)
		}

		merApiAuthRouter := v1.Group("/merchants/api").Use(middleware.StrictAuth(jwt, logger), middleware.AdminOnly())
		{
			merApiAuthRouter.POST("/create", merchantsApiHandler.CreateMerchantsApi)
			merApiAuthRouter.DELETE("/delete/:id", merchantsApiHandler.DeleteMerchantsApi)
			merApiAuthRouter.POST("/update", merchantsApiHandler.UpdateMerchantsApi)
			merApiAuthRouter.GET("/:id", merchantsApiHandler.GetMerchantsApi)
			merApiAuthRouter.POST("/list", merchantsApiHandler.GetMerchantsApiList)
		}

		// Add system wallet routes
		sysWalletGroup := v1.Group("/sys-wallet").Use(middleware.StrictAuth(jwt, logger), middleware.AdminOnly())
		{
			sysWalletGroup.GET("", sysWalletHandler.GetSysWallet)
			sysWalletGroup.POST("/create", sysWalletHandler.CreateSysWallet)
			sysWalletGroup.POST("/update", sysWalletHandler.UpdateSysWallet)
		}

		orderAuthRouter := v1.Group("/order").Use(middleware.StrictAuth(jwt, logger), middleware.AdminOnly())
		{
			orderAuthRouter.POST("/list", orderHandler.GetOrderList)
		}
		orderOutRouter := v1.Group("/order").Use(middleware.OrderAuth(logger))
		{
			orderOutRouter.POST("/create", orderHandler.CreateOrder)
			orderOutRouter.POST("/cancel", orderHandler.CancelOrder)
			orderOutRouter.POST("/success", orderHandler.SuccessOrder)
		}
		orderNoAuthRouter := v1.Group("/order").Use()
		{
			orderNoAuthRouter.GET("/order/:no", orderHandler.GetOrderPay)
			orderNoAuthRouter.POST("/:no/tx", orderHandler.GetOrderPayTx)
		}
		testAuthRouter := v1.Group("/").Use()
		{
			testAuthRouter.POST("/testcall", orderHandler.TestCall)
		}
		sysConfigRouter := v1.Group("/sys-config").Use(middleware.StrictAuth(jwt, logger), middleware.AdminOnly())
		{
			sysConfigRouter.GET("", sysConfigHandler.GetSysConfig)
			sysConfigRouter.POST("/create", sysConfigHandler.CreateSysConfig)
			sysConfigRouter.POST("/update", sysConfigHandler.UpdateSysConfig)
		}
	}

	return s
}
