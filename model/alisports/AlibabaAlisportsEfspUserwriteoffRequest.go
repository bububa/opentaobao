package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户核销 API请求
alibaba.alisports.efsp.userwriteoff

用户核销
*/
type AlibabaAlisportsEfspUserwriteoffRequest struct {
    model.Params
    // 订单编号
    _orderNo   string
    // 订单金额
    _sumAmount   int64
    // 健身房Id
    _gymId   string
    // 用户支付宝ID
    _alipayId   string
    // 补助金额
    _subsidyAmount   int64
}

// 初始化AlibabaAlisportsEfspUserwriteoffRequest对象
func NewAlibabaAlisportsEfspUserwriteoffRequest() *AlibabaAlisportsEfspUserwriteoffRequest{
    return &AlibabaAlisportsEfspUserwriteoffRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsEfspUserwriteoffRequest) GetApiMethodName() string {
    return "alibaba.alisports.efsp.userwriteoff"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsEfspUserwriteoffRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderNo Setter
// 订单编号
func (r *AlibabaAlisportsEfspUserwriteoffRequest) SetOrderNo(_orderNo string) error {
    r._orderNo = _orderNo
    r.Set("order_no", _orderNo)
    return nil
}

// OrderNo Getter
func (r AlibabaAlisportsEfspUserwriteoffRequest) GetOrderNo() string {
    return r._orderNo
}
// SumAmount Setter
// 订单金额
func (r *AlibabaAlisportsEfspUserwriteoffRequest) SetSumAmount(_sumAmount int64) error {
    r._sumAmount = _sumAmount
    r.Set("sum_amount", _sumAmount)
    return nil
}

// SumAmount Getter
func (r AlibabaAlisportsEfspUserwriteoffRequest) GetSumAmount() int64 {
    return r._sumAmount
}
// GymId Setter
// 健身房Id
func (r *AlibabaAlisportsEfspUserwriteoffRequest) SetGymId(_gymId string) error {
    r._gymId = _gymId
    r.Set("gym_id", _gymId)
    return nil
}

// GymId Getter
func (r AlibabaAlisportsEfspUserwriteoffRequest) GetGymId() string {
    return r._gymId
}
// AlipayId Setter
// 用户支付宝ID
func (r *AlibabaAlisportsEfspUserwriteoffRequest) SetAlipayId(_alipayId string) error {
    r._alipayId = _alipayId
    r.Set("alipay_id", _alipayId)
    return nil
}

// AlipayId Getter
func (r AlibabaAlisportsEfspUserwriteoffRequest) GetAlipayId() string {
    return r._alipayId
}
// SubsidyAmount Setter
// 补助金额
func (r *AlibabaAlisportsEfspUserwriteoffRequest) SetSubsidyAmount(_subsidyAmount int64) error {
    r._subsidyAmount = _subsidyAmount
    r.Set("subsidy_amount", _subsidyAmount)
    return nil
}

// SubsidyAmount Getter
func (r AlibabaAlisportsEfspUserwriteoffRequest) GetSubsidyAmount() int64 {
    return r._subsidyAmount
}