package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
外部合作商家重发电子凭证回调接口 API请求
taobao.vmarket.eticket.resend

外部合作商家重发电子凭证回调接口
*/
type TaobaoVmarketEticketResendRequest struct {
    model.Params
    // 订单编号
    orderId   int64
    // 重新发送的验证码及可验证次数的列表，多个码之间用英文逗号分割，需要包含此订单所有可用的码（如果订单总的有10个码，可用的是5个，那么这里设置的是5个可用的码）
    verifyCodes   string
    // 安全验证token,回传淘宝发通知时发过来的token串
    token   string
    // 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
    codemerchantId   int64
    // 不需要上传二维码图片的码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
    qrImages   string
}

// 初始化TaobaoVmarketEticketResendRequest对象
func NewTaobaoVmarketEticketResendRequest() *TaobaoVmarketEticketResendRequest{
    return &TaobaoVmarketEticketResendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketResendRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.resend"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketResendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单编号
func (r *TaobaoVmarketEticketResendRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoVmarketEticketResendRequest) GetOrderId() int64 {
    return r.orderId
}
// VerifyCodes Setter
// 重新发送的验证码及可验证次数的列表，多个码之间用英文逗号分割，需要包含此订单所有可用的码（如果订单总的有10个码，可用的是5个，那么这里设置的是5个可用的码）
func (r *TaobaoVmarketEticketResendRequest) SetVerifyCodes(verifyCodes string) error {
    r.verifyCodes = verifyCodes
    r.Set("verify_codes", verifyCodes)
    return nil
}

// VerifyCodes Getter
func (r TaobaoVmarketEticketResendRequest) GetVerifyCodes() string {
    return r.verifyCodes
}
// Token Setter
// 安全验证token,回传淘宝发通知时发过来的token串
func (r *TaobaoVmarketEticketResendRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoVmarketEticketResendRequest) GetToken() string {
    return r.token
}
// CodemerchantId Setter
// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
func (r *TaobaoVmarketEticketResendRequest) SetCodemerchantId(codemerchantId int64) error {
    r.codemerchantId = codemerchantId
    r.Set("codemerchant_id", codemerchantId)
    return nil
}

// CodemerchantId Getter
func (r TaobaoVmarketEticketResendRequest) GetCodemerchantId() int64 {
    return r.codemerchantId
}
// QrImages Setter
// 不需要上传二维码图片的码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
func (r *TaobaoVmarketEticketResendRequest) SetQrImages(qrImages string) error {
    r.qrImages = qrImages
    r.Set("qr_images", qrImages)
    return nil
}

// QrImages Getter
func (r TaobaoVmarketEticketResendRequest) GetQrImages() string {
    return r.qrImages
}
