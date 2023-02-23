package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenuidGetAPIRequest 获取授权账号对应的OpenUid API请求
// taobao.openuid.get
//
// 获取授权账号对应的OpenUid
type TaobaoOpenuidGetAPIRequest struct {
	model.Params
}

// NewTaobaoOpenuidGetRequest 初始化TaobaoOpenuidGetAPIRequest对象
func NewTaobaoOpenuidGetRequest() *TaobaoOpenuidGetAPIRequest {
	return &TaobaoOpenuidGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenuidGetAPIRequest) GetApiMethodName() string {
	return "taobao.openuid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenuidGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenuidGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
