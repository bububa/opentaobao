package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlifanyiMarketAuthenticateAPIRequest
第三方授权 API请求
alibaba.alifanyi.market.authenticate

第三方授权，获取授权码 */
type AlibabaAlifanyiMarketAuthenticateAPIRequest struct {
	model.Params
	// 有效时长
	_expireTime int64
}

// New
