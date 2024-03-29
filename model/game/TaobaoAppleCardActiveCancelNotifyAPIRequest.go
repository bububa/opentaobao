package game

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAppleCardActiveCancelNotifyAPIRequest 苹果卡密取消激活回调接口 API请求
// taobao.apple.card.active.cancel.notify
//
// 苹果卡密取消激活回调接口
type TaobaoAppleCardActiveCancelNotifyAPIRequest struct {
	model.Params
	// 卡信息
	_appleCardCancels []AppleCardCancelDto
	// 原商户申请激活时的订单号
	_orderNo string
	// 结果，000：整单取消激活成功  04： 取消激活失败（包括全部失败和部分失败，此时需以detail为准）
	_resultCode string
	// 网关订单号
	_gatewayOrderNo string
	// 描述
	_resultMsg string
}

// NewTaobaoAppleCardActiveCancelNotifyRequest 初始化TaobaoAppleCardActiveCancelNotifyAPIRequest对象
func NewTaobaoAppleCardActiveCancelNotifyRequest() *TaobaoAppleCardActiveCancelNotifyAPIRequest {
	return &TaobaoAppleCardActiveCancelNotifyAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAppleCardActiveCancelNotifyAPIRequest) Reset() {
	r._appleCardCancels = r._appleCardCancels[:0]
	r._orderNo = ""
	r._resultCode = ""
	r._gatewayOrderNo = ""
	r._resultMsg = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAppleCardActiveCancelNotifyAPIRequest) GetApiMethodName() string {
	return "taobao.apple.card.active.cancel.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAppleCardActiveCancelNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAppleCardActiveCancelNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppleCardCancels is AppleCardCancels Setter
// 卡信息
func (r *TaobaoAppleCardActiveCancelNotifyAPIRequest) SetAppleCardCancels(_appleCardCancels []AppleCardCancelDto) error {
	r._appleCardCancels = _appleCardCancels
	r.Set("apple_card_cancels", _appleCardCancels)
	return nil
}

// GetAppleCardCancels AppleCardCancels Getter
func (r TaobaoAppleCardActiveCancelNotifyAPIRequest) GetAppleCardCancels() []AppleCardCancelDto {
	return r._appleCardCancels
}

// SetOrderNo is OrderNo Setter
// 原商户申请激活时的订单号
func (r *TaobaoAppleCardActiveCancelNotifyAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r TaobaoAppleCardActiveCancelNotifyAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetResultCode is ResultCode Setter
// 结果，000：整单取消激活成功  04： 取消激活失败（包括全部失败和部分失败，此时需以detail为准）
func (r *TaobaoAppleCardActiveCancelNotifyAPIRequest) SetResultCode(_resultCode string) error {
	r._resultCode = _resultCode
	r.Set("result_code", _resultCode)
	return nil
}

// GetResultCode ResultCode Getter
func (r TaobaoAppleCardActiveCancelNotifyAPIRequest) GetResultCode() string {
	return r._resultCode
}

// SetGatewayOrderNo is GatewayOrderNo Setter
// 网关订单号
func (r *TaobaoAppleCardActiveCancelNotifyAPIRequest) SetGatewayOrderNo(_gatewayOrderNo string) error {
	r._gatewayOrderNo = _gatewayOrderNo
	r.Set("gateway_order_no", _gatewayOrderNo)
	return nil
}

// GetGatewayOrderNo GatewayOrderNo Getter
func (r TaobaoAppleCardActiveCancelNotifyAPIRequest) GetGatewayOrderNo() string {
	return r._gatewayOrderNo
}

// SetResultMsg is ResultMsg Setter
// 描述
func (r *TaobaoAppleCardActiveCancelNotifyAPIRequest) SetResultMsg(_resultMsg string) error {
	r._resultMsg = _resultMsg
	r.Set("result_msg", _resultMsg)
	return nil
}

// GetResultMsg ResultMsg Getter
func (r TaobaoAppleCardActiveCancelNotifyAPIRequest) GetResultMsg() string {
	return r._resultMsg
}

var poolTaobaoAppleCardActiveCancelNotifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAppleCardActiveCancelNotifyRequest()
	},
}

// GetTaobaoAppleCardActiveCancelNotifyRequest 从 sync.Pool 获取 TaobaoAppleCardActiveCancelNotifyAPIRequest
func GetTaobaoAppleCardActiveCancelNotifyAPIRequest() *TaobaoAppleCardActiveCancelNotifyAPIRequest {
	return poolTaobaoAppleCardActiveCancelNotifyAPIRequest.Get().(*TaobaoAppleCardActiveCancelNotifyAPIRequest)
}

// ReleaseTaobaoAppleCardActiveCancelNotifyAPIRequest 将 TaobaoAppleCardActiveCancelNotifyAPIRequest 放入 sync.Pool
func ReleaseTaobaoAppleCardActiveCancelNotifyAPIRequest(v *TaobaoAppleCardActiveCancelNotifyAPIRequest) {
	v.Reset()
	poolTaobaoAppleCardActiveCancelNotifyAPIRequest.Put(v)
}
