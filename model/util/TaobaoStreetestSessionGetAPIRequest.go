package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoStreetestSessionGetAPIRequest
根据获取压测用户的sessionKey API请求
taobao.streetest.session.get

根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测 */
type TaobaoStreetestSessionGetAPIRequest struct {
	model.Params
}

// New
