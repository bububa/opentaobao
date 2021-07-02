package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenTradeUserDeleteAPIRequest 删除奇门订单链路用户 API请求
// taobao.qimen.trade.user.delete
//
// 删除奇门订单链路用户
type TaobaoQimenTradeUserDeleteAPIRequest struct {
	model.Params
}

// NewTaobaoQimenTradeUserDeleteRequest 初始化TaobaoQimenTradeUserDeleteAPIRequest对象
func NewTaobaoQimenTradeUserDeleteRequest() *TaobaoQimenTradeUserDeleteAPIRequest {
	return &TaobaoQimenTradeUserDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenTradeUserDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.trade.user.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenTradeUserDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
