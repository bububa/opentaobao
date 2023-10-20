package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaostreetestsessiongetAPIRequest 根据获取压测用户的sessionKey API请求
// taobao.streetest.session.get
//
// 根据正常用户sessionKey获取对应压测用户的sessionKey，该sessionKey只能用户服务商全链路压测
type TaobaostreetestsessiongetAPIRequest struct {
	model.Params
}

// NewTaobaostreetestsessiongetRequest 初始化TaobaostreetestsessiongetAPIRequest对象
func NewTaobaostreetestsessiongetRequest() *TaobaostreetestsessiongetAPIRequest {
	return &TaobaostreetestsessiongetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaostreetestsessiongetAPIRequest) GetApiMethodName() string {
	return "taobao.streetest.session.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaostreetestsessiongetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaostreetestsessiongetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
