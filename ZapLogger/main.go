package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
)

var sugarLogger *zap.SugaredLogger

func main() {
	InitLogger()
	defer func(sugarLogger *zap.SugaredLogger) {
		err := sugarLogger.Sync()
		if err != nil {
			sugarLogger.Errorf("sync error")
		}
	}(sugarLogger)
	simpleHttpGet("www.baidu.com")
	simpleHttpGet("https://www.baidu.com")
}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter() zapcore.WriteSyncer {
	file, _ := os.Create("ZapLogger/log/test.log")
	return zapcore.AddSync(file)
}

func simpleHttpGet(url string) {
	sugarLogger.Debugf("Trying to hit GET request for %s", url)
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
	} else {
		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
		err := resp.Body.Close()
		if err != nil {
			return
		}
	}

}
