package user

import (
	"github.com/gin-gonic/gin"
	schemUser "www.miniton-gateway.com/app/schema/user"
	servUser "www.miniton-gateway.com/app/service/user"
	"www.miniton-gateway.com/utils/response"
)

func Detail(c *gin.Context) {
	var (
		err error
		req schemUser.DetailReq
	)
	tgUserInfo, err := schemUser.GetUserInfo(c.Request.Context())
	if err != nil {
		response.AuthErr(c)
		return
	}
	req = schemUser.DetailReq{}
	req.TGUserInfo = tgUserInfo
	resp, err := servUser.Detail(c.Request.Context(), &req)
	if err != nil {
		response.SysErr(c, err)
		return
	}
	response.Success(c, gin.H{"data": resp})
}
