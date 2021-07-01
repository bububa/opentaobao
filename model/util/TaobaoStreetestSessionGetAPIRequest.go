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

// NewTaobaoStreetestSessionGetRequest 初始化TaobaoStreetestSessionGetAPIRequest对象
func NewTaobaoStreetestSessionGetRequest() *TaobaoStreetestSessionGetAPIRequest {
	return &TaobaoStreetestSessionGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoStreetestSessionGetAPIRequest) GetApiMethodName() string {
	return "taobao.streetest.session.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoStreetestSessionGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
