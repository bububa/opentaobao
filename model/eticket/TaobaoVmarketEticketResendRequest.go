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
    _orderId   int64
    // 重新发送的验证码及可验证次数的列表，多个码之间用英文逗号分割，需要包含此订单所有可用的码（如果订单总的有10个码，可用的是5个，那么这里设置的是5个可用的码）
    _verifyCodes   string
    // 安全验证token,回传淘宝发通知时发过来的token串
    _token   string
    // 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
    _codemerchantId   int64
    // 不需要上传二维码图片的码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
    _qrImages   string
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
func (r *TaobaoVmarketEticketResendRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoVmarketEticketResendRequest) GetOrderId() int64 {
    return r._orderId
}
// VerifyCodes Setter
// 重新发送的验证码及可验证次数的列表，多个码之间用英文逗号分割，需要包含此订单所有可用的码（如果订单总的有10个码，可用的是5个，那么这里设置的是5个可用的码）
func (r *TaobaoVmarketEticketResendRequest) SetVerifyCodes(_verifyCodes string) error {
    r._verifyCodes = _verifyCodes
    r.Set("verify_codes", _verifyCodes)
    return nil
}

// VerifyCodes Getter
func (r TaobaoVmarketEticketResendRequest) GetVerifyCodes() string {
    return r._verifyCodes
}
// Token Setter
// 安全验证token,回传淘宝发通知时发过来的token串
func (r *TaobaoVmarketEticketResendRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoVmarketEticketResendRequest) GetToken() string {
    return r._token
}
// CodemerchantId Setter
// 码商ID，如果是码商，必须传，如果是信任卖家，不需要传
func (r *TaobaoVmarketEticketResendRequest) SetCodemerchantId(_codemerchantId int64) error {
    r._codemerchantId = _codemerchantId
    r.Set("codemerchant_id", _codemerchantId)
    return nil
}

// CodemerchantId Getter
func (r TaobaoVmarketEticketResendRequest) GetCodemerchantId() int64 {
    return r._codemerchantId
}
// QrImages Setter
// 不需要上传二维码图片的码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
func (r *TaobaoVmarketEticketResendRequest) SetQrImages(_qrImages string) error {
    r._qrImages = _qrImages
    r.Set("qr_images", _qrImages)
    return nil
}

// QrImages Getter
func (r TaobaoVmarketEticketResendRequest) GetQrImages() string {
    return r._qrImages
}
