package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	h "www.miniton-gateway.com/pkg/http"
	"www.miniton-gateway.com/pkg/log"
	"www.miniton-gateway.com/pkg/schema"
)

func Success(c *gin.Context, objs ...gin.H) {
	c.Set(schema.RespKey, objs)
	c.JSON(http.StatusOK, h.BuildSuccess(c, objs...))
}

func BindErr(c *gin.Context) {
	m := "request param valid"
	c.Set(schema.RespKey, m)
	h.Error(c, http.StatusBadRequest, m)
}

func ValidateErr(c *gin.Context, err error) {
	m := "request param valid err:[" + err.Error() + "]"
	c.Set(schema.RespKey, m)
	h.Error(c, http.StatusBadRequest, m)
}

func SysErr(c *gin.Context, err error) {
	c.Set(schema.RespKey, err.Error())
	h.Error(c, http.StatusInternalServerError, "system error")
}

func FoundErr(c *gin.Context) {
	e := NotFoundErr.Error()
	c.Set(schema.RespKey, e)
	h.Error(c, http.StatusNotFound, e)
}

func AuthErr(c *gin.Context) {
	e := NotFoundErr.Error()
	c.Set(schema.RespKey, e)
	h.Error(c, http.StatusBadRequest, e)
}

func Abort(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"errno":   code,
		"errmsg":  msg,
		"traceid": log.TraceID(ctx.Request.Context()),
	})
}
