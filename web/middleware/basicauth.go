// file: middleware/basicauth.go

package middleware

import "github.com/kataras/iris/v12/middleware/basicauth"

var users = map[string]string{
	"usr":   "pss",
	"admin": "admin",
}

var BasicAuth = basicauth.New(basicauth.Options{
	Allow:    basicauth.AllowUsers(users),
	Realm:    basicauth.DefaultRealm,
	MaxTries: 1,
})
