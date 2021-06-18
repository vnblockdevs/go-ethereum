package parser

import (
	"github.com/op/go-logging"
	"gopkg.in/natefinch/lumberjack.v2"
)

// var Block = logging.MustGetLogger("blocks")
// var Txs = logging.MustGetLogger("txs")
var Receipt = logging.MustGetLogger("txs")

func Init(dirpath string) {
	// blockPath := dirpath + "/blocks.log"
	// txPath := dirpath + "/txs.log"
	receiptPath := dirpath + "/receipts.log"
	// initLogger(blockPath, Block)
	// initLogger(txPath, Txs)
	initLogger(receiptPath, Receipt)
}

func initLogger(fileName string, logger *logging.Logger) {
	backend := createLogBackend(fileName, 1)
	backend1Leveled := logging.AddModuleLevel(backend)
	backend1Leveled.SetLevel(logging.INFO, "")
	logger.SetBackend(backend1Leveled)
}

func createLogBackend(name string, size int) logging.Backend {
	return logging.NewLogBackend(&lumberjack.Logger{
		Filename:   name,
		MaxSize:    size, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	}, "", 0)
}
