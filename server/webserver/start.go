package webserver

import (
	"net/http"

	"github.com/7sDream/rikka/common/logger"
)

var (
	isServeTLS bool
	l *logger.Logger
)

// StartRikkaWebServer start web server of rikka.
func StartRikkaWebServer(maxSizeByMb float64, argIsServeTLS bool, log *logger.Logger) string {

	if maxSizeByMb <= 0 {
		l.Fatal("Max file size can't be equal or less than 0, you set it to", maxSizeByMb)
	}

	isServeTLS = argIsServeTLS

	context.MaxSizeByMb = maxSizeByMb
	context.FavIconPath = FavIconTruePath

	l = log.SubLogger("[Web]")

	checkFiles()

	http.HandleFunc(RootPath, indexHandlerGenerator())
	http.HandleFunc(ViewPath, viewHandleGenerator())
	http.HandleFunc(StaticPath, staticFsHandlerGenerator())
	http.HandleFunc(FavIconOriginPath, favIconHandlerGenerator())

	l.Info("Rikka web server start successfully")

	return ViewPath
}
