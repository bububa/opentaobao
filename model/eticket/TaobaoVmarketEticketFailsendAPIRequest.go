package eticket

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketFailsendAPIRequest 无法发码回调 API请求
// taobao.vmarket.eticket.failsend
//
// 针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
type TaobaoVmarketEticketFailsendAPIRequest struct {
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

// NewTaobaoVmarketEticketFailsendRequest 初始化TaobaoVmarketEticketFailsendAPIRequest对象
func NewTaobaoVmarketEticketFailsendRequest() *TaobaoVmarketEticketFailsendAPIRequest {
	return &TaobaoVmarketEticketFailsendAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoVmarketEticketFailsendAPIRequest) Reset() {
	r._token = ""
	r._errorMsg = ""
	r._orderId = 0
	r._errorCode = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketFailsendAPIRequest) GetApiMethodName() string {
	return "taobao.vmarket.eticket.failsend"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketFailsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoVmarketEticketFailsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 发码通知时的token
func (r *TaobaoVmarketEticketFailsendAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoVmarketEticketFailsendAPIRequest) GetToken() string {
	return r._token
}

// SetErrorMsg is ErrorMsg Setter
// 错误信息
func (r *TaobaoVmarketEticketFailsendAPIRequest) SetErrorMsg(_errorMsg string) error {
	r._errorMsg = _errorMsg
	r.Set("error_msg", _errorMsg)
	return nil
}

// GetErrorMsg ErrorMsg Getter
func (r TaobaoVmarketEticketFailsendAPIRequest) GetErrorMsg() string {
	return r._errorMsg
}

// SetOrderId is OrderId Setter
// 订单号
func (r *TaobaoVmarketEticketFailsendAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoVmarketEticketFailsendAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetErrorCode is ErrorCode Setter
// 错误码
func (r *TaobaoVmarketEticketFailsendAPIRequest) SetErrorCode(_errorCode int64) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// GetErrorCode ErrorCode Getter
func (r TaobaoVmarketEticketFailsendAPIRequest) GetErrorCode() int64 {
	return r._errorCode
}

var poolTaobaoVmarketEticketFailsendAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoVmarketEticketFailsendRequest()
	},
}

// GetTaobaoVmarketEticketFailsendRequest 从 sync.Pool 获取 TaobaoVmarketEticketFailsendAPIRequest
func GetTaobaoVmarketEticketFailsendAPIRequest() *TaobaoVmarketEticketFailsendAPIRequest {
	return poolTaobaoVmarketEticketFailsendAPIRequest.Get().(*TaobaoVmarketEticketFailsendAPIRequest)
}

// ReleaseTaobaoVmarketEticketFailsendAPIRequest 将 TaobaoVmarketEticketFailsendAPIRequest 放入 sync.Pool
func ReleaseTaobaoVmarketEticketFailsendAPIRequest(v *TaobaoVmarketEticketFailsendAPIRequest) {
	v.Reset()
	poolTaobaoVmarketEticketFailsendAPIRequest.Put(v)
}
