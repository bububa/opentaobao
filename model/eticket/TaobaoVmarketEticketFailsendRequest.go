package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
无法发码回调 APIRequest
taobao.vmarket.eticket.failsend

针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
*/
type TaobaoVmarketEticketFailsendRequest struct {
    model.Params

    // 订单号
    orderId   int64 

    // 发码通知时的token
    token   string 

    // 错误码
    errorCode   int64 

    // 错误信息
    errorMsg   string 

}

func NewTaobaoVmarketEticketFailsendRequest() *TaobaoVmarketEticketFailsendRequest{
    return &TaobaoVmarketEticketFailsendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoVmarketEticketFailsendRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.failsend"
}

func (r TaobaoVmarketEticketFailsendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoVmarketEticketFailsendRequest) SetOrderId(orderId int64) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r TaobaoVmarketEticketFailsendRequest) GetOrderId() int64 {
    return r.orderId
}

func (r *TaobaoVmarketEticketFailsendRequest) SetToken(token string) error {
    r.token = token
    r.Set("token", token)
    return nil
}

func (r TaobaoVmarketEticketFailsendRequest) GetToken() string {
    return r.token
}

func (r *TaobaoVmarketEticketFailsendRequest) SetErrorCode(errorCode int64) error {
    r.errorCode = errorCode
    r.Set("error_code", errorCode)
    return nil
}

func (r TaobaoVmarketEticketFailsendRequest) GetErrorCode() int64 {
    return r.errorCode
}

func (r *TaobaoVmarketEticketFailsendRequest) SetErrorMsg(errorMsg string) error {
    r.errorMsg = errorMsg
    r.Set("error_msg", errorMsg)
    return nil
}

func (r TaobaoVmarketEticketFailsendRequest) GetErrorMsg() string {
    return r.errorMsg
}

