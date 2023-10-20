package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoservindustryfinancegeexorderupdateAPIRequest 即科订单结果更新回调 API请求
// taobao.servindustry.finance.geex.order.update
//
// 即科订单结果更新回调本地接口
type TaobaoservindustryfinancegeexorderupdateAPIRequest struct {
	model.Params
	// 原因
	_reason string
	// 订单号
	_alscOrderId string
	// 状态
	_applyStatus string
}

// NewTaobaoservindustryfinancegeexorderupdateRequest 初始化TaobaoservindustryfinancegeexorderupdateAPIRequest对象
func NewTaobaoservindustryfinancegeexorderupdateRequest() *TaobaoservindustryfinancegeexorderupdateAPIRequest {
	return &TaobaoservindustryfinancegeexorderupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoservindustryfinancegeexorderupdateAPIRequest) GetApiMethodName() string {
	return "taobao.servindustry.finance.geex.order.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoservindustryfinancegeexorderupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoservindustryfinancegeexorderupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReason is Reason Setter
// 原因
func (r *TaobaoservindustryfinancegeexorderupdateAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r TaobaoservindustryfinancegeexorderupdateAPIRequest) GetReason() string {
	return r._reason
}

// SetAlscOrderId is AlscOrderId Setter
// 订单号
func (r *TaobaoservindustryfinancegeexorderupdateAPIRequest) SetAlscOrderId(_alscOrderId string) error {
	r._alscOrderId = _alscOrderId
	r.Set("alsc_order_id", _alscOrderId)
	return nil
}

// GetAlscOrderId AlscOrderId Getter
func (r TaobaoservindustryfinancegeexorderupdateAPIRequest) GetAlscOrderId() string {
	return r._alscOrderId
}

// SetApplyStatus is ApplyStatus Setter
// 状态
func (r *TaobaoservindustryfinancegeexorderupdateAPIRequest) SetApplyStatus(_applyStatus string) error {
	r._applyStatus = _applyStatus
	r.Set("apply_status", _applyStatus)
	return nil
}

// GetApplyStatus ApplyStatus Getter
func (r TaobaoservindustryfinancegeexorderupdateAPIRequest) GetApplyStatus() string {
	return r._applyStatus
}
