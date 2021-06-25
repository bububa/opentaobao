package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子票券消费通知 APIRequest
taobao.vmarket.eticket.consume

外部合作商家电子票券消费回调接口
*/
type TaobaoVmarketEticketConsumeRequest struct {
    model.Params

    // 进行验码的电子凭证订单的订单ID
    orderId   int64 

    // 核销的码，只支持单个码，多个码核销需要多次调用
    verifyCode   string 

    // 核销份数
    consumeNum   int64 

    // 安全验证token,需要和发码通知中的token一致
    token   string 

    // 码商ID,是码商的话必须传递,如果是信任卖家不需要传
    codemerchantId   int64 

    // 机具ID(此参数信任卖家可不传递，码商必须传递)
    posid   string 

    // 手机后四位(没有特殊说明请不要传该参数)
    mobile   string 

    // 核销后需要重新生成的码，如果不需要重新生成码，不要传该参数
    newCode   string 

    // 自定义核销流水号，如果核销调用失败，可以用该核销流水号进行冲正操作，需要小于等于100个字符(a-zA-Z0-9_)；每次核销都是唯一的流水号
    serialNum   string 

    // 不需要上传二维码图片或者核销后不需重新生成码码商请不要传，需要传入二维码的码商请先调用taobao.vmarket.eticket.qrcode.upload接口，将返回的img_filename文件名称作为参数（如果二维码不变的话，也可将将发码时传入二维码文件名作为参数传入），文件名与参数new_code必须相互对应。
    qrImages   string 

}

func NewTaobaoVmarketEticketConsumeRequest() *TaobaoVmarketEticketConsumeRequest{
    return &TaobaoVmarketEticketConsumeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVmarketEticketConsumeRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.consume"
}

func (r TaobaoVmarketEticketConsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVmarketEticketConsumeRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoVmarketEticketConsumeRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TaobaoVmarketEticketConsumeRequest) SetVerifyCode(verifyCode string) error {
    r.verifyCode = verifyCode
    r.Set("verify_code", verifyCode)
    return nil
}

func (r TaobaoVmarketEticketConsumeRequest) GetVerifyCode() string {
    return r.verifyCode
}

func (r *TaobaoVmarketEticketConsumeRequest) SetConsumeNum(consumeNum int64) error {
    r.consumeNum = consumeNum
    r.Set("consume_num", consumeNum)
    return nil
}

func (r TaobaoVmarketEticketConsumeRequest) GetConsumeNum() int64 {
    return r.consumeNum
}

func (r *TaobaoVmarketEticketConsumeRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoVmarketEticketConsumeRequest) GetToken() string {
    return r.token
}

func (r *TaobaoVmarketEticketConsumeRequest) SetCodemerchantId(codemerchantId int64) error {
    r.codemerchantId = codemerchantId
    r.Set("codemerchant_id", codemerchantId)
    return nil
}

func (r TaobaoVmarketEticketConsumeRequest) GetCodemerchantId() int64 {
    return r.codemerchantId
}

func (r *TaobaoVmarketEticketConsumeRequest) SetPosid(posid string) error {
    r.posid = posid
    r.Set("posid", posid)
    return nil
}

func (r TaobaoVmarketEticketConsumeRequest) GetPosid() string {
    return r.posid
}

func (r *TaobaoVmarketEticketConsumeRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r TaobaoVmarketEticketConsumeRequest) GetMobile() string {
    return r.mobile
}

func (r *TaobaoVmarketEticketConsumeRequest) SetNewCode(newCode string) error {
    r.newCode = newCode
    r.Set("new_code", newCode)
    return nil
}

func (r TaobaoVmarketEticketConsumeRequest) GetNewCode() string {
    return r.newCode
}

func (r *TaobaoVmarketEticketConsumeRequest) SetSerialNum(serialNum string) error {
    r.serialNum = serialNum
    r.Set("serial_num", serialNum)
    return nil
}

func (r TaobaoVmarketEticketConsumeRequest) GetSerialNum() string {
    return r.serialNum
}

func (r *TaobaoVmarketEticketConsumeRequest) SetQrImages(qrImages string) error {
    r.qrImages = qrImages
    r.Set("qr_images", qrImages)
    return nil
}

func (r TaobaoVmarketEticketConsumeRequest) GetQrImages() string {
    return r.qrImages
}

