package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户详细信息 API请求
alibaba.alisports.efsp.getuserinfo

阿里体育-体育健身-获取用户详细信息
*/
type AlibabaAlisportsEfspGetuserinfoRequest struct {
    model.Params
    // 支付宝ID
    alipayId   string
}

// 初始化AlibabaAlisportsEfspGetuserinfoRequest对象
func NewAlibabaAlisportsEfspGetuserinfoRequest() *AlibabaAlisportsEfspGetuserinfoRequest{
    return &AlibabaAlisportsEfspGetuserinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlisportsEfspGetuserinfoRequest) GetApiMethodName() string {
    return "alibaba.alisports.efsp.getuserinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlisportsEfspGetuserinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayId Setter
// 支付宝ID
func (r *AlibabaAlisportsEfspGetuserinfoRequest) SetAlipayId(alipayId string) error {
    r.alipayId = alipayId
    r.Set("alipay_id", alipayId)
    return nil
}

// AlipayId Getter
func (r AlibabaAlisportsEfspGetuserinfoRequest) GetAlipayId() string {
    return r.alipayId
}
