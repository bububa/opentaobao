package game

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAppleCardActiveApplyNotifyAPIRequest 苹果卡密申请激活回调接口 API请求
// taobao.apple.card.active.apply.notify
//
// 苹果卡密申请激活回调接口
type TaobaoAppleCardActiveApplyNotifyAPIRequest struct {
	model.Params
	// 卡列表
	_appleCards []AppleCardDto
	// 网关订单号
	_gatewayOrderNo string
	// 描述
	_resultMsg string
	// 商户唯一订单号
	_orderNo string
	// 结果，000：成功，其他皆为错误 04： 订单处理失败(商户可退款，其他不可退款)
	_resultCode string
}

// NewTaobaoAppleCardActiveApplyNotifyRequest 初始化TaobaoAppleCardActiveApplyNotifyAPIRequest对象
func NewTaobaoAppleCardActiveApplyNotifyRequest() *TaobaoAppleCardActiveApplyNotifyAPIRequest {
	return &TaobaoAppleCardActiveApplyNotifyAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAppleCardActiveApplyNotifyAPIRequest) Reset() {
	r._appleCards = r._appleCards[:0]
	r._gatewayOrderNo = ""
	r._resultMsg = ""
	r._orderNo = ""
	r._resultCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAppleCardActiveApplyNotifyAPIRequest) GetApiMethodName() string {
	return "taobao.apple.card.active.apply.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAppleCardActiveApplyNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAppleCardActiveApplyNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppleCards is AppleCards Setter
// 卡列表
func (r *TaobaoAppleCardActiveApplyNotifyAPIRequest) SetAppleCards(_appleCards []AppleCardDto) error {
	r._appleCards = _appleCards
	r.Set("apple_cards", _appleCards)
	return nil
}

// GetAppleCards AppleCards Getter
func (r TaobaoAppleCardActiveApplyNotifyAPIRequest) GetAppleCards() []AppleCardDto {
	return r._appleCards
}

// SetGatewayOrderNo is GatewayOrderNo Setter
// 网关订单号
func (r *TaobaoAppleCardActiveApplyNotifyAPIRequest) SetGatewayOrderNo(_gatewayOrderNo string) error {
	r._gatewayOrderNo = _gatewayOrderNo
	r.Set("gateway_order_no", _gatewayOrderNo)
	return nil
}

// GetGatewayOrderNo GatewayOrderNo Getter
func (r TaobaoAppleCardActiveApplyNotifyAPIRequest) GetGatewayOrderNo() string {
	return r._gatewayOrderNo
}

// SetResultMsg is ResultMsg Setter
// 描述
func (r *TaobaoAppleCardActiveApplyNotifyAPIRequest) SetResultMsg(_resultMsg string) error {
	r._resultMsg = _resultMsg
	r.Set("result_msg", _resultMsg)
	return nil
}

// GetResultMsg ResultMsg Getter
func (r TaobaoAppleCardActiveApplyNotifyAPIRequest) GetResultMsg() string {
	return r._resultMsg
}

// SetOrderNo is OrderNo Setter
// 商户唯一订单号
func (r *TaobaoAppleCardActiveApplyNotifyAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r TaobaoAppleCardActiveApplyNotifyAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetResultCode is ResultCode Setter
// 结果，000：成功，其他皆为错误 04： 订单处理失败(商户可退款，其他不可退款)
func (r *TaobaoAppleCardActiveApplyNotifyAPIRequest) SetResultCode(_resultCode string) error {
	r._resultCode = _resultCode
	r.Set("result_code", _resultCode)
	return nil
}

// GetResultCode ResultCode Getter
func (r TaobaoAppleCardActiveApplyNotifyAPIRequest) GetResultCode() string {
	return r._resultCode
}

var poolTaobaoAppleCardActiveApplyNotifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAppleCardActiveApplyNotifyRequest()
	},
}

// GetTaobaoAppleCardActiveApplyNotifyRequest 从 sync.Pool 获取 TaobaoAppleCardActiveApplyNotifyAPIRequest
func GetTaobaoAppleCardActiveApplyNotifyAPIRequest() *TaobaoAppleCardActiveApplyNotifyAPIRequest {
	return poolTaobaoAppleCardActiveApplyNotifyAPIRequest.Get().(*TaobaoAppleCardActiveApplyNotifyAPIRequest)
}

// ReleaseTaobaoAppleCardActiveApplyNotifyAPIRequest 将 TaobaoAppleCardActiveApplyNotifyAPIRequest 放入 sync.Pool
func ReleaseTaobaoAppleCardActiveApplyNotifyAPIRequest(v *TaobaoAppleCardActiveApplyNotifyAPIRequest) {
	v.Reset()
	poolTaobaoAppleCardActiveApplyNotifyAPIRequest.Put(v)
}
