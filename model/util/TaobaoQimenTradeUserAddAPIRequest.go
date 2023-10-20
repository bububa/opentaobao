package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimentradeuseraddAPIRequest 添加奇门订单链路用户 API请求
// taobao.qimen.trade.user.add
//
// 添加奇门订单链路用户
type TaobaoqimentradeuseraddAPIRequest struct {
	model.Params
	// 商家备注
	_memo string
}

// NewTaobaoqimentradeuseraddRequest 初始化TaobaoqimentradeuseraddAPIRequest对象
func NewTaobaoqimentradeuseraddRequest() *TaobaoqimentradeuseraddAPIRequest {
	return &TaobaoqimentradeuseraddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimentradeuseraddAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.trade.user.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimentradeuseraddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimentradeuseraddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMemo is Memo Setter
// 商家备注
func (r *TaobaoqimentradeuseraddAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TaobaoqimentradeuseraddAPIRequest) GetMemo() string {
	return r._memo
}
