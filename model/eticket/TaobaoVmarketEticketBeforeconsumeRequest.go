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
    _orderId   int64
    // 需要验的码
    _verifyCode   string
    // 安全验证token，需要和发码通知中的token一致
    _token   string
    // 码商ID,是码商的话必须传递,如果是信任卖家不需要传
    _codemerchantId   int64
    // 操作员身份ID，如果是码商必须传,如果是信任卖家不需要传
    _posid   string
    // 手机号码后四位,没有特殊说明请不要传
    _mobile   string
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
func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoVmarketEticketBeforeconsumeRequest) GetOrderId() int64 {
    return r._orderId
}
// VerifyCode Setter
// 需要验的码
func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetVerifyCode(_verifyCode string) error {
    r._verifyCode = _verifyCode
    r.Set("verify_code", _verifyCode)
    return nil
}

// VerifyCode Getter
func (r TaobaoVmarketEticketBeforeconsumeRequest) GetVerifyCode() string {
    return r._verifyCode
}
// Token Setter
// 安全验证token，需要和发码通知中的token一致
func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoVmarketEticketBeforeconsumeRequest) GetToken() string {
    return r._token
}
// CodemerchantId Setter
// 码商ID,是码商的话必须传递,如果是信任卖家不需要传
func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetCodemerchantId(_codemerchantId int64) error {
    r._codemerchantId = _codemerchantId
    r.Set("codemerchant_id", _codemerchantId)
    return nil
}

// CodemerchantId Getter
func (r TaobaoVmarketEticketBeforeconsumeRequest) GetCodemerchantId() int64 {
    return r._codemerchantId
}
// Posid Setter
// 操作员身份ID，如果是码商必须传,如果是信任卖家不需要传
func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetPosid(_posid string) error {
    r._posid = _posid
    r.Set("posid", _posid)
    return nil
}

// Posid Getter
func (r TaobaoVmarketEticketBeforeconsumeRequest) GetPosid() string {
    return r._posid
}
// Mobile Setter
// 手机号码后四位,没有特殊说明请不要传
func (r *TaobaoVmarketEticketBeforeconsumeRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r TaobaoVmarketEticketBeforeconsumeRequest) GetMobile() string {
    return r._mobile
}
