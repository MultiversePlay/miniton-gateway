package area

import (
	"github.com/gin-gonic/gin"
	schemArea "www.miniton-gateway.com/app/schema/area"
	"www.miniton-gateway.com/utils/response"
)

func CountryList(c *gin.Context) {
	response.Success(c, gin.H{"data": schemArea.CountryBuild})
}
