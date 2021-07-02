package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsEfspUserplaceorderAPIRequest 用户完成支付同步订单 API请求
// alibaba.alisports.efsp.userplaceorder
//
// 用户完成支付同步订单
type AlibabaAlisportsEfspUserplaceorderAPIRequest struct {
	model.Params
	// 青橙订单的json
	_orderJson string
}

// NewAlibabaAlisportsEfspUserplaceorderRequest 初始化AlibabaAlisportsEfspUserplaceorderAPIRequest对象
func NewAlibabaAlisportsEfspUserplaceorderRequest() *AlibabaAlisportsEfspUserplaceorderAPIRequest {
	return &AlibabaAlisportsEfspUserplaceorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsEfspUserplaceorderAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.efsp.userplaceorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsEfspUserplaceorderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderJson Setter
// 青橙订单的json
func (r *AlibabaAlisportsEfspUserplaceorderAPIRequest) SetOrderJson(_orderJson string) error {
	r._orderJson = _orderJson
	r.Set("order_json", _orderJson)
	return nil
}

// Get OrderJson Getter
func (r AlibabaAlisportsEfspUserplaceorderAPIRequest) GetOrderJson() string {
	return r._orderJson
}
