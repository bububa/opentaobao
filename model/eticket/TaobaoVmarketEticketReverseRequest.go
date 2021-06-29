package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证冲正接口 API请求
taobao.vmarket.eticket.reverse

电子凭证平台冲正接口
*/
type TaobaoVmarketEticketReverseRequest struct {
    model.Params
    // 进行验码的电子凭证订单的订单ID
    orderId   int64
    // 冲正的码，只支持单个码
    reverseCode   string
    // 冲正份数（必须是和被冲正的核销记录的份数一致）
    reverseNum   int64
    // 需要冲正的核销记录对应核销流水号（对应的核销操作时候传递的自定义流水号）
    consumeSecialNum   string
    // 所有冲正后需要重新生成的码和对应的次数。码和次数之间用英文冒号分隔，多个码之间用英文逗号分隔。如果冲正后不需要重新生成码，留空
    verifyCodes   string
    // 不需要上传二维码图片或者冲正后不需要变更码的请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数，多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
    qrImages   string
    // 安全验证token，需要和该订单发码通知中的token一致
    token   string
    // 码商ID，是码商的话必须传递，如果是信任卖家不要传
    codemerchantId   int64
    // 机具id，如果是码商必须传，如果是信任卖家不要传
    posid   string
}

// 初始化TaobaoVmarketEticketReverseRequest对象
func NewTaobaoVmarketEticketReverseRequest() *TaobaoVmarketEticketReverseRequest{
    return &TaobaoVmarketEticketReverseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketReverseRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.reverse"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketReverseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 进行验码的电子凭证订单的订单ID
func (r *TaobaoVmarketEticketReverseRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoVmarketEticketReverseRequest) GetOrderId() int64 {
    return r.orderId
}
// ReverseCode Setter
// 冲正的码，只支持单个码
func (r *TaobaoVmarketEticketReverseRequest) SetReverseCode(reverseCode string) error {
    r.reverseCode = reverseCode
    r.Set("reverse_code", reverseCode)
    return nil
}

// ReverseCode Getter
func (r TaobaoVmarketEticketReverseRequest) GetReverseCode() string {
    return r.reverseCode
}
// ReverseNum Setter
// 冲正份数（必须是和被冲正的核销记录的份数一致）
func (r *TaobaoVmarketEticketReverseRequest) SetReverseNum(reverseNum int64) error {
    r.reverseNum = reverseNum
    r.Set("reverse_num", reverseNum)
    return nil
}

// ReverseNum Getter
func (r TaobaoVmarketEticketReverseRequest) GetReverseNum() int64 {
    return r.reverseNum
}
// ConsumeSecialNum Setter
// 需要冲正的核销记录对应核销流水号（对应的核销操作时候传递的自定义流水号）
func (r *TaobaoVmarketEticketReverseRequest) SetConsumeSecialNum(consumeSecialNum string) error {
    r.consumeSecialNum = consumeSecialNum
    r.Set("consume_secial_num", consumeSecialNum)
    return nil
}

// ConsumeSecialNum Getter
func (r TaobaoVmarketEticketReverseRequest) GetConsumeSecialNum() string {
    return r.consumeSecialNum
}
// VerifyCodes Setter
// 所有冲正后需要重新生成的码和对应的次数。码和次数之间用英文冒号分隔，多个码之间用英文逗号分隔。如果冲正后不需要重新生成码，留空
func (r *TaobaoVmarketEticketReverseRequest) SetVerifyCodes(verifyCodes string) error {
    r.verifyCodes = verifyCodes
    r.Set("verify_codes", verifyCodes)
    return nil
}

// VerifyCodes Getter
func (r TaobaoVmarketEticketReverseRequest) GetVerifyCodes() string {
    return r.verifyCodes
}
// QrImages Setter
// 不需要上传二维码图片或者冲正后不需要变更码的请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数，多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
func (r *TaobaoVmarketEticketReverseRequest) SetQrImages(qrImages string) error {
    r.qrImages = qrImages
    r.Set("qr_images", qrImages)
    return nil
}

// QrImages Getter
func (r TaobaoVmarketEticketReverseRequest) GetQrImages() string {
    return r.qrImages
}
// Token Setter
// 安全验证token，需要和该订单发码通知中的token一致
func (r *TaobaoVmarketEticketReverseRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoVmarketEticketReverseRequest) GetToken() string {
    return r.token
}
// CodemerchantId Setter
// 码商ID，是码商的话必须传递，如果是信任卖家不要传
func (r *TaobaoVmarketEticketReverseRequest) SetCodemerchantId(codemerchantId int64) error {
    r.codemerchantId = codemerchantId
    r.Set("codemerchant_id", codemerchantId)
    return nil
}

// CodemerchantId Getter
func (r TaobaoVmarketEticketReverseRequest) GetCodemerchantId() int64 {
    return r.codemerchantId
}
// Posid Setter
// 机具id，如果是码商必须传，如果是信任卖家不要传
func (r *TaobaoVmarketEticketReverseRequest) SetPosid(posid string) error {
    r.posid = posid
    r.Set("posid", posid)
    return nil
}

// Posid Getter
func (r TaobaoVmarketEticketReverseRequest) GetPosid() string {
    return r.posid
}
