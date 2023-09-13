package auth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	schemUser "www.miniton-gateway.com/app/schema/user"
	"www.miniton-gateway.com/pkg/config"
	"www.miniton-gateway.com/utils/response"
)

type (
	WebAppUser struct {
	}
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get(schemUser.HeaderAuthKey)
		if auth == "" {
			response.Abort(c, http.StatusBadRequest, "header auth empty")
			return
		}
		token := config.Config.TgConfig.Token
		userInfo, err := verifyAuth(auth, token)
		if err != nil {
			response.Abort(c, http.StatusBadRequest, err.Error())
			return
		}

		ctx := context.WithValue(c.Request.Context(), schemUser.TGUserKey{}, userInfo)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func verifyAuth(auth string, token string) (userInfo *schemUser.TGUserInfo, err error) {
	initData, err := url.ParseQuery(auth)
	if err != nil {
		return
	}
	dataToCheck := []string{}
	for k, v := range initData {
		if k == "hash" {
			continue
		}

		dataToCheck = append(dataToCheck, fmt.Sprintf("%s=%s", k, v[0]))
	}

	sort.Strings(dataToCheck)

	secret := hmac.New(sha256.New, []byte("WebAppData"))
	secret.Write([]byte(token))

	hHash := hmac.New(sha256.New, secret.Sum(nil))
	hHash.Write([]byte(strings.Join(dataToCheck, "\n")))

	hash := hex.EncodeToString(hHash.Sum(nil))

	if initData.Get("hash") != hash {
		err = errors.New("auth error")
		return
	}

	userInfo = new(schemUser.TGUserInfo)
	err = json.Unmarshal([]byte(initData.Get("user")), &userInfo)
	if err != nil {
		err = errors.New("auth user error")
		return
	}
	return
}
