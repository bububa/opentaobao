package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaomodifyaddressopenAPIRequest 淘宝自助修改地址服务开通 API请求
// taobao.modifyaddress.open
//
// 商家自助修改地址服务开通
type TaobaomodifyaddressopenAPIRequest struct {
	model.Params
}

// NewTaobaomodifyaddressopenRequest 初始化TaobaomodifyaddressopenAPIRequest对象
func NewTaobaomodifyaddressopenRequest() *TaobaomodifyaddressopenAPIRequest {
	return &TaobaomodifyaddressopenAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaomodifyaddressopenAPIRequest) GetApiMethodName() string {
	return "taobao.modifyaddress.open"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaomodifyaddressopenAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaomodifyaddressopenAPIRequest) GetRawParams() model.Params {
	return r.Params
}
