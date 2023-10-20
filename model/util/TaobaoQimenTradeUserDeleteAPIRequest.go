package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimentradeuserdeleteAPIRequest 删除奇门订单链路用户 API请求
// taobao.qimen.trade.user.delete
//
// 删除奇门订单链路用户
type TaobaoqimentradeuserdeleteAPIRequest struct {
	model.Params
}

// NewTaobaoqimentradeuserdeleteRequest 初始化TaobaoqimentradeuserdeleteAPIRequest对象
func NewTaobaoqimentradeuserdeleteRequest() *TaobaoqimentradeuserdeleteAPIRequest {
	return &TaobaoqimentradeuserdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimentradeuserdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.trade.user.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimentradeuserdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimentradeuserdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}
