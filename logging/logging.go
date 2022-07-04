package logging

import (
	"io"
	"log"
	"os"
)

var (
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	WarningLogger *log.Logger
	BlocksLogger  *log.Logger
)

func SetupLogger() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		ErrorLogger.Fatal(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, file)

	InfoLogger = log.New(multiWriter, "INFO:		", log.LstdFlags|log.Lshortfile|log.LUTC|log.Lmicroseconds)
	WarningLogger = log.New(multiWriter, "WARNING:	", log.LstdFlags|log.Lshortfile|log.LUTC|log.Lmicroseconds)
	ErrorLogger = log.New(multiWriter, "ERROR:		", log.LstdFlags|log.Lshortfile|log.LUTC|log.Lmicroseconds)
	BlocksLogger = log.New(multiWriter, "BLOCKS:		", log.LstdFlags|log.Lshortfile|log.LUTC|log.Lmicroseconds)
}
