package interfaces

import "github.com/gin-gonic/gin"

type IHealthcheckServices interface {
	HealtcheckService() (string, error)
}

type IHealthcheckRepository interface {
}

type IHealthcheckHandler interface {
	HealtcheckHandlerHTTP(c *gin.Context)
}
