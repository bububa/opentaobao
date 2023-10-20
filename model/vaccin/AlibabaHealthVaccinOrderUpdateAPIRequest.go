package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthvaccinorderupdateAPIRequest 回传取号状态 API请求
// alibaba.health.vaccin.order.update
//
// 回传取号状态
type AlibabahealthvaccinorderupdateAPIRequest struct {
	model.Params
	// 操作类型
	_actionType string
	// 预约单id
	_orderId string
	// 未接种原因
	_isvNotInoculateReason string
	// 时间戳
	_actionTime int64
}

// NewAlibabahealthvaccinorderupdateRequest 初始化AlibabahealthvaccinorderupdateAPIRequest对象
func NewAlibabahealthvaccinorderupdateRequest() *AlibabahealthvaccinorderupdateAPIRequest {
	return &AlibabahealthvaccinorderupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahealthvaccinorderupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.order.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahealthvaccinorderupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahealthvaccinorderupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActionType is ActionType Setter
// 操作类型
func (r *AlibabahealthvaccinorderupdateAPIRequest) SetActionType(_actionType string) error {
	r._actionType = _actionType
	r.Set("action_type", _actionType)
	return nil
}

// GetActionType ActionType Getter
func (r AlibabahealthvaccinorderupdateAPIRequest) GetActionType() string {
	return r._actionType
}

// SetOrderId is OrderId Setter
// 预约单id
func (r *AlibabahealthvaccinorderupdateAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahealthvaccinorderupdateAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetIsvNotInoculateReason is IsvNotInoculateReason Setter
// 未接种原因
func (r *AlibabahealthvaccinorderupdateAPIRequest) SetIsvNotInoculateReason(_isvNotInoculateReason string) error {
	r._isvNotInoculateReason = _isvNotInoculateReason
	r.Set("isv_not_inoculate_reason", _isvNotInoculateReason)
	return nil
}

// GetIsvNotInoculateReason IsvNotInoculateReason Getter
func (r AlibabahealthvaccinorderupdateAPIRequest) GetIsvNotInoculateReason() string {
	return r._isvNotInoculateReason
}

// SetActionTime is ActionTime Setter
// 时间戳
func (r *AlibabahealthvaccinorderupdateAPIRequest) SetActionTime(_actionTime int64) error {
	r._actionTime = _actionTime
	r.Set("action_time", _actionTime)
	return nil
}

// GetActionTime ActionTime Getter
func (r AlibabahealthvaccinorderupdateAPIRequest) GetActionTime() int64 {
	return r._actionTime
}
