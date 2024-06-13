package internals

import "net/http"

type FindrAppService interface{
	Register() http.HandlerFunc
	Login() http.HandlerFunc
	Logout() http.HandlerFunc
}

