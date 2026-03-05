package main

import (
	"log"
	"log/slog"

	"github.com/go-uber/zap"
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/ladev74/linter/internal/analyzer"
	"github.com/ladev74/linter/internal/config"
)

// TODO: в README написать почему не вынес в конфиг поддерживаемые логгеры (в тз жестко прописаны какие логгеры нужно поддерживать)
// TODO: указать почему в тестах нет zap

func main() {
	cfg, err := config.New("config/example.yaml")
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	singlechecker.Main(analyzer.New(&cfg.Analyzer))

	logger := zap.Logger{}

	logger.Info("StartIng linter")
	slog.Info("StartIng linter")
	logger.Info("хых")
	logger.Info("server started!🚀")
	logger.Info("connection failed!!!")
	logger.Info("warning: something went wrong...")
	logger.Info("passive pass: password:")

	//path := ""
	//zapPath := "zap"
	//slogPath := "slog"
	//
	//fmt.Println(path == slogPath) || (path == zapPath)
}
