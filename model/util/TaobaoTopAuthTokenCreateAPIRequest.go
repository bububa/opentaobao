package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTopAuthTokenCreateAPIRequest
获取Access Token API请求
taobao.top.auth.token.create

用户通过code换获取access_token，https only */
type TaobaoTopAuthTokenCreateAPIRequest struct {
	model.Params
	// 授权code，grantType==authorization_code 时需要
	_code string
	// 与生成code的uuid配对
	_uuid string
}

// New
