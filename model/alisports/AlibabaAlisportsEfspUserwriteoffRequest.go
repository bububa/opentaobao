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
    orderNo   string
    // 订单金额
    sumAmount   int64
    // 健身房Id
    gymId   string
    // 用户支付宝ID
    alipayId   string
    // 补助金额
    subsidyAmount   int64
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
func (r *AlibabaAlisportsEfspUserwriteoffRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

// OrderNo Getter
func (r AlibabaAlisportsEfspUserwriteoffRequest) GetOrderNo() string {
    return r.orderNo
}
// SumAmount Setter
// 订单金额
func (r *AlibabaAlisportsEfspUserwriteoffRequest) SetSumAmount(sumAmount int64) error {
    r.sumAmount = sumAmount
    r.Set("sum_amount", sumAmount)
    return nil
}

// SumAmount Getter
func (r AlibabaAlisportsEfspUserwriteoffRequest) GetSumAmount() int64 {
    return r.sumAmount
}
// GymId Setter
// 健身房Id
func (r *AlibabaAlisportsEfspUserwriteoffRequest) SetGymId(gymId string) error {
    r.gymId = gymId
    r.Set("gym_id", gymId)
    return nil
}

// GymId Getter
func (r AlibabaAlisportsEfspUserwriteoffRequest) GetGymId() string {
    return r.gymId
}
// AlipayId Setter
// 用户支付宝ID
func (r *AlibabaAlisportsEfspUserwriteoffRequest) SetAlipayId(alipayId string) error {
    r.alipayId = alipayId
    r.Set("alipay_id", alipayId)
    return nil
}

// AlipayId Getter
func (r AlibabaAlisportsEfspUserwriteoffRequest) GetAlipayId() string {
    return r.alipayId
}
// SubsidyAmount Setter
// 补助金额
func (r *AlibabaAlisportsEfspUserwriteoffRequest) SetSubsidyAmount(subsidyAmount int64) error {
    r.subsidyAmount = subsidyAmount
    r.Set("subsidy_amount", subsidyAmount)
    return nil
}

// SubsidyAmount Getter
func (r AlibabaAlisportsEfspUserwriteoffRequest) GetSubsidyAmount() int64 {
    return r.subsidyAmount
}
