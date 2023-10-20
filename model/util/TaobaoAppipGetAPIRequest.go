package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAppipGetAPIRequest 获取ISV发起请求服务器IP API请求
// taobao.appip.get
//
// 获取ISV发起请求服务器IP
type TaobaoAppipGetAPIRequest struct {
	model.Params
}

// NewTaobaoAppipGetRequest 初始化TaobaoAppipGetAPIRequest对象
func NewTaobaoAppipGetRequest() *TaobaoAppipGetAPIRequest {
	return &TaobaoAppipGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAppipGetAPIRequest) GetApiMethodName() string {
	return "taobao.appip.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAppipGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAppipGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
