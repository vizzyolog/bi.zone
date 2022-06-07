package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"app/interfaces/agentConsumer"
	"app/interfaces/mongodb"
	"app/interfaces/rest"
	"app/interfaces/web"

	"app/pkg/cryptor"
	"app/pkg/graceful"
	"app/pkg/logger"
	"app/pkg/middlewares"

	"app/usecases/events"
	"app/usecases/users"
)

func main() {

	cfg := loadConfig()
	lg := logger.New(os.Stderr, cfg.LogLevel, cfg.LogFormat)

	db, closeSession, err := mongodb.Connect(cfg.MongoURI, true)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}
	defer closeSession()

	log.Println("setting up rest api service")

	userStore := mongodb.NewUserStore(db)
	eventStore := mongodb.NewEventStore(db)

	userRegistration := users.NewRegistrar(lg, userStore)
	userRetriever := users.NewRetriever(lg, userStore)

	eventPersist := events.NewPersistar(lg, eventStore)
	eventRetriver := events.NewRetriever(lg, eventStore)

	restHandler := rest.New(lg, userRegistration, userRetriever, eventRetriver)
	webHandler, err := web.New(lg, web.Config{
		TemplateDir: cfg.TemplateDir,
		StaticDir:   cfg.StaticDir,
	})
	if err != nil {
		lg.Fatalf("failed to setup web handler: %v", err)
	}

	srv := setupServer(cfg, lg, webHandler, restHandler)
	lg.Infof("listening for requests on :8080...")
	if err := srv.ListenAndServe(); err != nil {
		lg.Fatalf("http server exited: %s", err)
	}

	cryptor := cryptor.New()

	eventConsumer := agentConsumer.New(cryptor, eventPersist, eventStore)

}

func setupServer(cfg config, lg logger.Logger, web http.Handler, rest http.Handler) *graceful.Server {
	rest = middlewares.WithBasicAuth(lg, rest,
		middlewares.UserVerifierFunc(func(ctx context.Context, name, secret string) bool {
			return secret == "secret@123"
		}),
	)

	router := mux.NewRouter()
	router.PathPrefix("/api").Handler(http.StripPrefix("/api", rest))
	router.PathPrefix("/").Handler(web)

	handler := middlewares.WithRequestLogging(lg, router)
	handler = middlewares.WithRecovery(lg, handler)

	srv := graceful.NewServer(handler, cfg.GracefulTimeout, os.Interrupt)
	srv.Log = lg.Errorf
	srv.Addr = cfg.Addr
	return srv
}

func setupConsumer(cfg config, lg logger.Logger, eventStore *mongodb.EventStore) *graceful.Server {

	return
}

type config struct {
	Addr            string
	LogLevel        string
	LogFormat       string
	StaticDir       string
	TemplateDir     string
	GracefulTimeout time.Duration
	MongoURI        string
}

func loadConfig() config {
	viper.SetDefault("MONGO_URI", "mongodb://localhost/Bi.zone")
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("LOG_FORMAT", "text")
	viper.SetDefault("ADDR", ":8080")
	viper.SetDefault("STATIC_DIR", "./web/static/")
	viper.SetDefault("TEMPLATE_DIR", "./web/templates/")
	viper.SetDefault("GRACEFUL_TIMEOUT", 20*time.Second)

	viper.ReadInConfig()
	viper.AutomaticEnv()

	return config{
		// application configuration
		Addr:            viper.GetString("ADDR"),
		StaticDir:       viper.GetString("STATIC_DIR"),
		TemplateDir:     viper.GetString("TEMPLATE_DIR"),
		LogLevel:        viper.GetString("LOG_LEVEL"),
		LogFormat:       viper.GetString("LOG_FORMAT"),
		GracefulTimeout: viper.GetDuration("GRACEFUL_TIMEOUT"),

		// store configuration
		MongoURI: viper.GetString("MONGO_URI"),
	}
}

func parseTime(s string) time.Duration {
	t, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return 0
	}

	return time.Duration(t) * time.Second

}
