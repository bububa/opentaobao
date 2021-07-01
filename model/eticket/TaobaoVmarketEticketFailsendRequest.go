package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
无法发码回调 API请求
taobao.vmarket.eticket.failsend

针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
*/
type TaobaoVmarketEticketFailsendAPIRequest struct {
    model.Params
    // 订单号
    _orderId   int64
    // 发码通知时的token
    _token   string
    // 错误码
    _errorCode   int64
    // 错误信息
    _errorMsg   string
}

// 初始化TaobaoVmarketEticketFailsendAPIRequest对象
func NewTaobaoVmarketEticketFailsendRequest() *TaobaoVmarketEticketFailsendAPIRequest{
    return &TaobaoVmarketEticketFailsendAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketFailsendAPIRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.failsend"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketFailsendAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单号
func (r *TaobaoVmarketEticketFailsendAPIRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoVmarketEticketFailsendAPIRequest) GetOrderId() int64 {
    return r._orderId
}
// Token Setter
// 发码通知时的token
func (r *TaobaoVmarketEticketFailsendAPIRequest) SetToken(_token string) error {
    r._token = _token
    r.Set("token", _token)
    return nil
}

// Token Getter
func (r TaobaoVmarketEticketFailsendAPIRequest) GetToken() string {
    return r._token
}
// ErrorCode Setter
// 错误码
func (r *TaobaoVmarketEticketFailsendAPIRequest) SetErrorCode(_errorCode int64) error {
    r._errorCode = _errorCode
    r.Set("error_code", _errorCode)
    return nil
}

// ErrorCode Getter
func (r TaobaoVmarketEticketFailsendAPIRequest) GetErrorCode() int64 {
    return r._errorCode
}
// ErrorMsg Setter
// 错误信息
func (r *TaobaoVmarketEticketFailsendAPIRequest) SetErrorMsg(_errorMsg string) error {
    r._errorMsg = _errorMsg
    r.Set("error_msg", _errorMsg)
    return nil
}

// ErrorMsg Getter
func (r TaobaoVmarketEticketFailsendAPIRequest) GetErrorMsg() string {
    return r._errorMsg
}
