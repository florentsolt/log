package log

import (
	"net/http"

	"github.com/rs/zerolog/hlog"
)

func HttpHandler() func(http.Handler) http.Handler {
	return hlog.NewHandler(instance.parent)
}
