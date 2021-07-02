package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsEfspUsercancelorderAPIRequest 用户取消订单 API请求
// alibaba.alisports.efsp.usercancelorder
//
// 用户取消订单
type AlibabaAlisportsEfspUsercancelorderAPIRequest struct {
	model.Params
	// 订单编号
	_orderNo string
	// 用户支付宝ID
	_alipayId string
}

// NewAlibabaAlisportsEfspUsercancelorderRequest 初始化AlibabaAlisportsEfspUsercancelorderAPIRequest对象
func NewAlibabaAlisportsEfspUsercancelorderRequest() *AlibabaAlisportsEfspUsercancelorderAPIRequest {
	return &AlibabaAlisportsEfspUsercancelorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsEfspUsercancelorderAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.efsp.usercancelorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsEfspUsercancelorderAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOrderNo is OrderNo Setter
// 订单编号
func (r *AlibabaAlisportsEfspUsercancelorderAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r AlibabaAlisportsEfspUsercancelorderAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetAlipayId is AlipayId Setter
// 用户支付宝ID
func (r *AlibabaAlisportsEfspUsercancelorderAPIRequest) SetAlipayId(_alipayId string) error {
	r._alipayId = _alipayId
	r.Set("alipay_id", _alipayId)
	return nil
}

// GetAlipayId AlipayId Getter
func (r AlibabaAlisportsEfspUsercancelorderAPIRequest) GetAlipayId() string {
	return r._alipayId
}
