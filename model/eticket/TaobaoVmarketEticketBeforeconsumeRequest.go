package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证验码前置确认 APIRequest
taobao.vmarket.eticket.beforeconsume

商家验码之前的调用接口，用来同步到最新的订单状态并判断是否可以进行验码操作
*/
type TaobaoVmarketEticketBeforeconsumeRequest struct {
    model.Params

    // 需要验码的电子凭证订单ID
    orderId   int64 

    // 需要验的码
    verifyCode   string 

    // 安全验证token，需要和发码通知中的token一致
    token   string 

    // 码商ID,是码商的话必须传递,如果是信任卖家不需要传
    codemerchantId   int64 

    // 操作员身份ID，如果是码商必须传,如果是信任卖家不需要传
    posid   string 

    // 手机号码后四位,没有特殊说明请不要传
    mobile   string 

}

func NewTaobaoVmarketEticketBeforeconsumeRequest() *TaobaoVmarketEticketBeforeconsumeRequest{
    return &TaobaoVmarketEticketBeforeconsumeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVmarketEticketBeforeconsumeRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.beforeconsume"
}

func (r TaobaoVmarketEticketBeforeconsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoVmarketEticketBeforeconsumeRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetVerifyCode(verifyCode string) error {
    r.verifyCode = verifyCode
    r.Set("verify_code", verifyCode)
    return nil
}

func (r TaobaoVmarketEticketBeforeconsumeRequest) GetVerifyCode() string {
    return r.verifyCode
}

func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoVmarketEticketBeforeconsumeRequest) GetToken() string {
    return r.token
}

func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetCodemerchantId(codemerchantId int64) error {
    r.codemerchantId = codemerchantId
    r.Set("codemerchant_id", codemerchantId)
    return nil
}

func (r TaobaoVmarketEticketBeforeconsumeRequest) GetCodemerchantId() int64 {
    return r.codemerchantId
}

func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetPosid(posid string) error {
    r.posid = posid
    r.Set("posid", posid)
    return nil
}

func (r TaobaoVmarketEticketBeforeconsumeRequest) GetPosid() string {
    return r.posid
}

func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r TaobaoVmarketEticketBeforeconsumeRequest) GetMobile() string {
    return r.mobile
}

