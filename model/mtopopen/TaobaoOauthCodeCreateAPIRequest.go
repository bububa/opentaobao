package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOauthCodeCreateAPIRequest
淘宝OauthCode颁发 API请求
taobao.oauth.code.create

手淘无线开放的oauthCode颁发接口 */
type TaobaoOauthCodeCreateAPIRequest struct {
	model.Params
	// mock param
	_test int64
}

// New
