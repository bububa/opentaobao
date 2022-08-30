package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJipiaoAgentOrderDetailAPIRequest 【机票代理商订单】订单详情 API请求
// taobao.jipiao.agent.order.detail
//
// 根据淘宝系统订单号获取订单详情信息
type TaobaoJipiaoAgentOrderDetailAPIRequest struct {
	model.Params
}

// NewTaobaoJipiaoAgentOrderDetailRequest 初始化TaobaoJipiaoAgentOrderDetailAPIRequest对象
func NewTaobaoJipiaoAgentOrderDetailRequest() *TaobaoJipiaoAgentOrderDetailAPIRequest {
	return &TaobaoJipiaoAgentOrderDetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJipiaoAgentOrderDetailAPIRequest) GetApiMethodName() string {
	return "taobao.jipiao.agent.order.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJipiaoAgentOrderDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
