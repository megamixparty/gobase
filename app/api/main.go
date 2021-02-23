package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/joeshaw/envdecode"
	"github.com/megamixparty/gobase/config"
	"github.com/megamixparty/gobase/lib/handler"
	"github.com/megamixparty/gobase/lib/repository/mysql"
	rr "github.com/megamixparty/gobase/lib/repository/redis"
	"github.com/megamixparty/gobase/lib/service/user_service"
	"github.com/subosito/gotenv"
)

func main() {
	var (
		err         error
		envConfig   config.EnvConfig
		mysqlDB     *sqlx.DB
		redisClient *redis.Client
	)
	gotenv.Load(".env")

	envdecode.Decode(&envConfig)
	mysqlDB, err = config.NewMysql(&envConfig)
	if err != nil {
		panic(err)
	}

	redisClient, err = config.NewRedis(&envConfig)
	if err != nil {
		panic(err)
	}

	// Repository definition
	mysqlRepo := mysql.NewMysqlRepository(mysqlDB)
	redisRepo := rr.NewRedisRepository(redisClient)

	// Service definition
	userService := user_service.NewUserService(mysqlRepo, redisRepo)

	// Handler definition
	userHandler := handler.NewUserHandler(userService)

	// Routes definition
	r := chi.NewRouter()
	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetUsers)
	})

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", envConfig.Port),
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Service is available at %s\n", s.Addr)
	log.Fatal(s.ListenAndServe())
}
