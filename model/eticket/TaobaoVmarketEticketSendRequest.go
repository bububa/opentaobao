package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家电子凭证发码成功回调接口 APIRequest
taobao.vmarket.eticket.send

外部商家成功发码回调接口
*/
type TaobaoVmarketEticketSendRequest struct {
    model.Params

    // 订单编号
    orderId   int64 

    // 发送成功的验证码及可验证次数的列表，码和可验证次数用英文冒号分隔，多个码之间用英文逗号分隔，所有字符都为英文半角
    verifyCodes   string 

    // 安全验证token，需要和发码通知中的token一致
    token   string 

    // 码商ID,是码商的话必须传递,如果是信任卖家,不需要传
    codemerchantId   int64 

    // 不需要上传二维码图片的码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数，多个文件名用逗号隔开且与参数verify_codes按从左到有的顺序一一对应。
    qrImages   string 

}

func NewTaobaoVmarketEticketSendRequest() *TaobaoVmarketEticketSendRequest{
    return &TaobaoVmarketEticketSendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVmarketEticketSendRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.send"
}

func (r TaobaoVmarketEticketSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVmarketEticketSendRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoVmarketEticketSendRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TaobaoVmarketEticketSendRequest) SetVerifyCodes(verifyCodes string) error {
    r.verifyCodes = verifyCodes
    r.Set("verify_codes", verifyCodes)
    return nil
}

func (r TaobaoVmarketEticketSendRequest) GetVerifyCodes() string {
    return r.verifyCodes
}

func (r *TaobaoVmarketEticketSendRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoVmarketEticketSendRequest) GetToken() string {
    return r.token
}

func (r *TaobaoVmarketEticketSendRequest) SetCodemerchantId(codemerchantId int64) error {
    r.codemerchantId = codemerchantId
    r.Set("codemerchant_id", codemerchantId)
    return nil
}

func (r TaobaoVmarketEticketSendRequest) GetCodemerchantId() int64 {
    return r.codemerchantId
}

func (r *TaobaoVmarketEticketSendRequest) SetQrImages(qrImages string) error {
    r.qrImages = qrImages
    r.Set("qr_images", qrImages)
    return nil
}

func (r TaobaoVmarketEticketSendRequest) GetQrImages() string {
    return r.qrImages
}

