package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户取消订单 APIRequest
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

func NewAlibabaAlisportsEfspUsercancelorderRequest() *AlibabaAlisportsEfspUsercancelorderRequest{
    return &AlibabaAlisportsEfspUsercancelorderRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsEfspUsercancelorderRequest) GetApiMethodName() string {
    return "alibaba.alisports.efsp.usercancelorder"
}

func (r AlibabaAlisportsEfspUsercancelorderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsEfspUsercancelorderRequest) SetOrderNo(orderNo string) error {
    r.orderNo = orderNo
    r.Set("order_no", orderNo)
    return nil
}

func (r AlibabaAlisportsEfspUsercancelorderRequest) GetOrderNo() string {
    return r.orderNo
}

func (r *AlibabaAlisportsEfspUsercancelorderRequest) SetAlipayId(alipayId string) error {
    r.alipayId = alipayId
    r.Set("alipay_id", alipayId)
    return nil
}

func (r AlibabaAlisportsEfspUsercancelorderRequest) GetAlipayId() string {
    return r.alipayId
}

