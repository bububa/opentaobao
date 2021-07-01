package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
扫码识别会员接口 API请求
alibaba.wdk.member.qrcode.identify

根据用户输入的付款码（支付宝、盒马、淘宝）、商家等信息，查询当前用户的基本信息及对应会员卡信息
*/
type AlibabaWdkMemberQrcodeIdentifyAPIRequest struct {
    model.Params
    // 付款码
    _qrCode   string
}

// 初始化AlibabaWdkMemberQrcodeIdentifyAPIRequest对象
func NewAlibabaWdkMemberQrcodeIdentifyRequest() *AlibabaWdkMemberQrcodeIdentifyAPIRequest{
    return &AlibabaWdkMemberQrcodeIdentifyAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkMemberQrcodeIdentifyAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.member.qrcode.identify"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkMemberQrcodeIdentifyAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QrCode Setter
// 付款码
func (r *AlibabaWdkMemberQrcodeIdentifyAPIRequest) SetQrCode(_qrCode string) error {
    r._qrCode = _qrCode
    r.Set("qr_code", _qrCode)
    return nil
}

// QrCode Getter
func (r AlibabaWdkMemberQrcodeIdentifyAPIRequest) GetQrCode() string {
    return r._qrCode
}
