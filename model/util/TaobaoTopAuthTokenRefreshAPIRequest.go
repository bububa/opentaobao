package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTopAuthTokenRefreshAPIRequest
刷新Access Token API请求
taobao.top.auth.token.refresh

根据refresh_token重新生成token，目前只有服务市场订购类应用可以刷新token，其他类型应用（如商家后台）使用固定时长token，不提供刷新功能。 */
type TaobaoTopAuthTokenRefreshAPIRequest struct {
	model.Params
	// grantType==refresh_token 时需要
	_refreshToken string
}

// New
