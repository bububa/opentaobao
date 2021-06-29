package alisports

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户详细信息 APIRequest
alibaba.alisports.efsp.getuserinfo

阿里体育-体育健身-获取用户详细信息
*/
type AlibabaAlisportsEfspGetuserinfoRequest struct {
    model.Params

    // 支付宝ID
    alipayId   string 

}

func NewAlibabaAlisportsEfspGetuserinfoRequest() *AlibabaAlisportsEfspGetuserinfoRequest{
    return &AlibabaAlisportsEfspGetuserinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlisportsEfspGetuserinfoRequest) GetApiMethodName() string {
    return "alibaba.alisports.efsp.getuserinfo"
}

func (r AlibabaAlisportsEfspGetuserinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlisportsEfspGetuserinfoRequest) SetAlipayId(alipayId string) error {
    r.alipayId = alipayId
    r.Set("alipay_id", alipayId)
    return nil
}

func (r AlibabaAlisportsEfspGetuserinfoRequest) GetAlipayId() string {
    return r.alipayId
}

