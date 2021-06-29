package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户取消订单 API请求
alibaba.alisports.efsp.usercancelorder

用户取消订单
*/
type AlibabaAlisportsEfspUsercancelorderRequest struct {
    model.Params
    // 订单编号
    orderNo   string
    // 用户支付宝ID
    alipayId   string
}

// 初始化AlibabaAlisportsEfspUsercancelorderRequest对象
func NewAlibabaAlisportsEfspUsercancelorderRequest() *AlibabaAlisportsEfspUsercancelorderRequest{
    return &AlibabaAlisportsEfspUsercancelorderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsEfspUsercancelorderRequest) GetApiMethodName() string {
    return "alibaba.alisports.efsp.usercancelorder"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsEfspUsercancelorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderNo Setter
// 订单编号
func (r *AlibabaAlisportsEfspUsercancelorderRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

// OrderNo Getter
func (r AlibabaAlisportsEfspUsercancelorderRequest) GetOrderNo() string {
    return r.orderNo
}
// AlipayId Setter
// 用户支付宝ID
func (r *AlibabaAlisportsEfspUsercancelorderRequest) SetAlipayId(alipayId string) error {
    r.alipayId = alipayId
    r.Set("alipay_id", alipayId)
    return nil
}

// AlipayId Getter
func (r AlibabaAlisportsEfspUsercancelorderRequest) GetAlipayId() string {
    return r.alipayId
}
