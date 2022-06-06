package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"app/interfaces/mongodb"
	"app/interfaces/rest"
	"app/pkg/cryptor"
	"app/pkg/logger"

	"github.com/spf13/viper"
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

	eventStore := mongodb.NewEventStore(db)
	userStore := mongodb.NewUserStore(db)

	eventRetriver := event.NewRetriever(lg, eventStore)

	restHandler := rest.New(lg, userRegistration, userRetriever, postRet, postPub)
	webHandler, err := web.New(lg, web.Config{
		TemplateDir: cfg.TemplateDir,
		StaticDir:   cfg.StaticDir,
	})
	if err != nil {
		lg.Fatalf("failed to setup web handler: %v", err)
	}

	cryptor := cryptor.Cryptor.New()
	_ = cryptor

	tcpSrv := tcpSrv.NewTC
	_ = tcpSrv

	restHandler := rest.New(lg, userRegistration, userRetriever, postRet, postPub)
	webHandler, err := web.New(lg, web.Config{
		TemplateDir: cfg.TemplateDir,
		StaticDir:   cfg.StaticDir,
	})
	if err != nil {
		lg.Fatalf("failed to setup web handler: %v", err)
	}

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
