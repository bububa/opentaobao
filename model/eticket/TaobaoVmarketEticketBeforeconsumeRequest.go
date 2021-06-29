package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子凭证验码前置确认 API请求
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

// 初始化TaobaoVmarketEticketBeforeconsumeRequest对象
func NewTaobaoVmarketEticketBeforeconsumeRequest() *TaobaoVmarketEticketBeforeconsumeRequest{
    return &TaobaoVmarketEticketBeforeconsumeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketBeforeconsumeRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.beforeconsume"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketBeforeconsumeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 需要验码的电子凭证订单ID
func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r TaobaoVmarketEticketBeforeconsumeRequest) GetOrderId() int64 {
    return r.orderId
}
// VerifyCode Setter
// 需要验的码
func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetVerifyCode(verifyCode string) error {
    r.verifyCode = verifyCode
    r.Set("verify_code", verifyCode)
    return nil
}

// VerifyCode Getter
func (r TaobaoVmarketEticketBeforeconsumeRequest) GetVerifyCode() string {
    return r.verifyCode
}
// Token Setter
// 安全验证token，需要和发码通知中的token一致
func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

// Token Getter
func (r TaobaoVmarketEticketBeforeconsumeRequest) GetToken() string {
    return r.token
}
// CodemerchantId Setter
// 码商ID,是码商的话必须传递,如果是信任卖家不需要传
func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetCodemerchantId(codemerchantId int64) error {
    r.codemerchantId = codemerchantId
    r.Set("codemerchant_id", codemerchantId)
    return nil
}

// CodemerchantId Getter
func (r TaobaoVmarketEticketBeforeconsumeRequest) GetCodemerchantId() int64 {
    return r.codemerchantId
}
// Posid Setter
// 操作员身份ID，如果是码商必须传,如果是信任卖家不需要传
func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetPosid(posid string) error {
    r.posid = posid
    r.Set("posid", posid)
    return nil
}

// Posid Getter
func (r TaobaoVmarketEticketBeforeconsumeRequest) GetPosid() string {
    return r.posid
}
// Mobile Setter
// 手机号码后四位,没有特殊说明请不要传
func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r TaobaoVmarketEticketBeforeconsumeRequest) GetMobile() string {
    return r.mobile
}
