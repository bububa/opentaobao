package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoocapcontractsignedgetAPIRequest 用户是否签署支付宝代扣协议 API请求
// taobao.oc.ap.contractsigned.get
//
// 用户是否签署支付宝代扣协议
type TaobaoocapcontractsignedgetAPIRequest struct {
	model.Params
}

// NewTaobaoocapcontractsignedgetRequest 初始化TaobaoocapcontractsignedgetAPIRequest对象
func NewTaobaoocapcontractsignedgetRequest() *TaobaoocapcontractsignedgetAPIRequest {
	return &TaobaoocapcontractsignedgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoocapcontractsignedgetAPIRequest) GetApiMethodName() string {
	return "taobao.oc.ap.contractsigned.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoocapcontractsignedgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoocapcontractsignedgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
