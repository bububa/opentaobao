package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlisportsEfspUserwriteoffAPIRequest 用户核销 API请求
// alibaba.alisports.efsp.userwriteoff
//
// 用户核销
type AlibabaAlisportsEfspUserwriteoffAPIRequest struct {
	model.Params
	// 订单编号
	_orderNo string
	// 订单金额
	_sumAmount int64
	// 健身房Id
	_gymId string
	// 用户支付宝ID
	_alipayId string
	// 补助金额
	_subsidyAmount int64
}

// NewAlibabaAlisportsEfspUserwriteoffRequest 初始化AlibabaAlisportsEfspUserwriteoffAPIRequest对象
func NewAlibabaAlisportsEfspUserwriteoffRequest() *AlibabaAlisportsEfspUserwriteoffAPIRequest {
	return &AlibabaAlisportsEfspUserwriteoffAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsEfspUserwriteoffAPIRequest) GetApiMethodName() string {
	return "alibaba.alisports.efsp.userwriteoff"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsEfspUserwriteoffAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderNo Setter
// 订单编号
func (r *AlibabaAlisportsEfspUserwriteoffAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// Get OrderNo Getter
func (r AlibabaAlisportsEfspUserwriteoffAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// Set is SumAmount Setter
// 订单金额
func (r *AlibabaAlisportsEfspUserwriteoffAPIRequest) SetSumAmount(_sumAmount int64) error {
	r._sumAmount = _sumAmount
	r.Set("sum_amount", _sumAmount)
	return nil
}

// Get SumAmount Getter
func (r AlibabaAlisportsEfspUserwriteoffAPIRequest) GetSumAmount() int64 {
	return r._sumAmount
}

// Set is GymId Setter
// 健身房Id
func (r *AlibabaAlisportsEfspUserwriteoffAPIRequest) SetGymId(_gymId string) error {
	r._gymId = _gymId
	r.Set("gym_id", _gymId)
	return nil
}

// Get GymId Getter
func (r AlibabaAlisportsEfspUserwriteoffAPIRequest) GetGymId() string {
	return r._gymId
}

// Set is AlipayId Setter
// 用户支付宝ID
func (r *AlibabaAlisportsEfspUserwriteoffAPIRequest) SetAlipayId(_alipayId string) error {
	r._alipayId = _alipayId
	r.Set("alipay_id", _alipayId)
	return nil
}

// Get AlipayId Getter
func (r AlibabaAlisportsEfspUserwriteoffAPIRequest) GetAlipayId() string {
	return r._alipayId
}

// Set is SubsidyAmount Setter
// 补助金额
func (r *AlibabaAlisportsEfspUserwriteoffAPIRequest) SetSubsidyAmount(_subsidyAmount int64) error {
	r._subsidyAmount = _subsidyAmount
	r.Set("subsidy_amount", _subsidyAmount)
	return nil
}

// Get SubsidyAmount Getter
func (r AlibabaAlisportsEfspUserwriteoffAPIRequest) GetSubsidyAmount() int64 {
	return r._subsidyAmount
}
