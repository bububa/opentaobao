package game

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoapplecardactiveapplynotifyAPIRequest 苹果卡密申请激活回调接口 API请求
// taobao.apple.card.active.apply.notify
//
// 苹果卡密申请激活回调接口
type TaobaoapplecardactiveapplynotifyAPIRequest struct {
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

// NewTaobaoapplecardactiveapplynotifyRequest 初始化TaobaoapplecardactiveapplynotifyAPIRequest对象
func NewTaobaoapplecardactiveapplynotifyRequest() *TaobaoapplecardactiveapplynotifyAPIRequest {
	return &TaobaoapplecardactiveapplynotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoapplecardactiveapplynotifyAPIRequest) GetApiMethodName() string {
	return "taobao.apple.card.active.apply.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoapplecardactiveapplynotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoapplecardactiveapplynotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppleCards is AppleCards Setter
// 卡列表
func (r *TaobaoapplecardactiveapplynotifyAPIRequest) SetAppleCards(_appleCards []AppleCardDto) error {
	r._appleCards = _appleCards
	r.Set("apple_cards", _appleCards)
	return nil
}

// GetAppleCards AppleCards Getter
func (r TaobaoapplecardactiveapplynotifyAPIRequest) GetAppleCards() []AppleCardDto {
	return r._appleCards
}

// SetGatewayOrderNo is GatewayOrderNo Setter
// 网关订单号
func (r *TaobaoapplecardactiveapplynotifyAPIRequest) SetGatewayOrderNo(_gatewayOrderNo string) error {
	r._gatewayOrderNo = _gatewayOrderNo
	r.Set("gateway_order_no", _gatewayOrderNo)
	return nil
}

// GetGatewayOrderNo GatewayOrderNo Getter
func (r TaobaoapplecardactiveapplynotifyAPIRequest) GetGatewayOrderNo() string {
	return r._gatewayOrderNo
}

// SetResultMsg is ResultMsg Setter
// 描述
func (r *TaobaoapplecardactiveapplynotifyAPIRequest) SetResultMsg(_resultMsg string) error {
	r._resultMsg = _resultMsg
	r.Set("result_msg", _resultMsg)
	return nil
}

// GetResultMsg ResultMsg Getter
func (r TaobaoapplecardactiveapplynotifyAPIRequest) GetResultMsg() string {
	return r._resultMsg
}

// SetOrderNo is OrderNo Setter
// 商户唯一订单号
func (r *TaobaoapplecardactiveapplynotifyAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r TaobaoapplecardactiveapplynotifyAPIRequest) GetOrderNo() string {
	return r._orderNo
}

// SetResultCode is ResultCode Setter
// 结果，000：成功，其他皆为错误 04： 订单处理失败(商户可退款，其他不可退款)
func (r *TaobaoapplecardactiveapplynotifyAPIRequest) SetResultCode(_resultCode string) error {
	r._resultCode = _resultCode
	r.Set("result_code", _resultCode)
	return nil
}

// GetResultCode ResultCode Getter
func (r TaobaoapplecardactiveapplynotifyAPIRequest) GetResultCode() string {
	return r._resultCode
}
