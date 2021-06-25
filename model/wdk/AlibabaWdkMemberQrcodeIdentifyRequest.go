package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
扫码识别会员接口 APIRequest
alibaba.wdk.member.qrcode.identify

根据用户输入的付款码（支付宝、盒马、淘宝）、商家等信息，查询当前用户的基本信息及对应会员卡信息
*/
type AlibabaWdkMemberQrcodeIdentifyRequest struct {
    model.Params

    // 付款码
    qrCode   string 

}

func NewAlibabaWdkMemberQrcodeIdentifyRequest() *AlibabaWdkMemberQrcodeIdentifyRequest{
    return &AlibabaWdkMemberQrcodeIdentifyRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkMemberQrcodeIdentifyRequest) GetApiMethodName() string {
    return "alibaba.wdk.member.qrcode.identify"
}

func (r AlibabaWdkMemberQrcodeIdentifyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkMemberQrcodeIdentifyRequest) SetQrCode(qrCode string) error {
    r.qrCode = qrCode
    r.Set("qr_code", qrCode)
    return nil
}

func (r AlibabaWdkMemberQrcodeIdentifyRequest) GetQrCode() string {
    return r.qrCode
}

