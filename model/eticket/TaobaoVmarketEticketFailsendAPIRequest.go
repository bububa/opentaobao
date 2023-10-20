package eticket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaovmarketeticketfailsendAPIRequest 无法发码回调 API请求
// taobao.vmarket.eticket.failsend
//
// 针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
type TaobaovmarketeticketfailsendAPIRequest struct {
	model.Params
	// 发码通知时的token
	_token string
	// 错误信息
	_errorMsg string
	// 订单号
	_orderId int64
	// 错误码
	_errorCode int64
}

// NewTaobaovmarketeticketfailsendRequest 初始化TaobaovmarketeticketfailsendAPIRequest对象
func NewTaobaovmarketeticketfailsendRequest() *TaobaovmarketeticketfailsendAPIRequest {
	return &TaobaovmarketeticketfailsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaovmarketeticketfailsendAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.failsend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaovmarketeticketfailsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaovmarketeticketfailsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 发码通知时的token
func (r *TaobaovmarketeticketfailsendAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaovmarketeticketfailsendAPIRequest) GetToken() string {
	return r._token
}

// SetErrorMsg is ErrorMsg Setter
// 错误信息
func (r *TaobaovmarketeticketfailsendAPIRequest) SetErrorMsg(_errorMsg string) error {
	r._errorMsg = _errorMsg
	r.Set("error_msg", _errorMsg)
	return nil
}

// GetErrorMsg ErrorMsg Getter
func (r TaobaovmarketeticketfailsendAPIRequest) GetErrorMsg() string {
	return r._errorMsg
}

// SetOrderId is OrderId Setter
// 订单号
func (r *TaobaovmarketeticketfailsendAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaovmarketeticketfailsendAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetErrorCode is ErrorCode Setter
// 错误码
func (r *TaobaovmarketeticketfailsendAPIRequest) SetErrorCode(_errorCode int64) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// GetErrorCode ErrorCode Getter
func (r TaobaovmarketeticketfailsendAPIRequest) GetErrorCode() int64 {
	return r._errorCode
}
