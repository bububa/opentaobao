package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAppipGetAPIRequest
获取ISV发起请求服务器IP API请求
taobao.appip.get

获取ISV发起请求服务器IP */
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
func (r TaobaoAppipGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
