package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/joeshaw/envdecode"
	"github.com/megamixparty/gobase/config"
	"github.com/megamixparty/gobase/lib/handler"
	"github.com/megamixparty/gobase/lib/repository/mysql"
	"github.com/megamixparty/gobase/lib/service/user_service"
	"github.com/subosito/gotenv"
)

func main() {
	var (
		err       error
		envConfig config.EnvConfig
		mysqlDB   *sqlx.DB
	)
	gotenv.Load(".env")

	envdecode.Decode(&envConfig)
	mysqlDB, err = config.NewMysql(&envConfig)
	if err != nil {
		panic(err)
	}

	// Repository definition
	userRepo := mysql.NewMysqlRepository(mysqlDB)

	// Service definition
	userService := user_service.NewUserService(userRepo)

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
