package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoappipgetAPIRequest 获取ISV发起请求服务器IP API请求
// taobao.appip.get
//
// 获取ISV发起请求服务器IP
type TaobaoappipgetAPIRequest struct {
	model.Params
}

// NewTaobaoappipgetRequest 初始化TaobaoappipgetAPIRequest对象
func NewTaobaoappipgetRequest() *TaobaoappipgetAPIRequest {
	return &TaobaoappipgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoappipgetAPIRequest) GetApiMethodName() string {
	return "taobao.appip.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoappipgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoappipgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
