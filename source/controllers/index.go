package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexController function
func IndexController(context *gin.Context) {
	context.Redirect(http.StatusMovedPermanently, "https://bitbucket.org/MattiasPernhult/friends-api/overview")
}
