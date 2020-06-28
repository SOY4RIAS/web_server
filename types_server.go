package web_server

import (
	"net/http"
)

type MiddleWare func(http.HandlerFunc) http.HandlerFunc
