package vaccin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHealthVaccinOrderUpdateAPIRequest 回传取号状态 API请求
// alibaba.health.vaccin.order.update
//
// 回传取号状态
type AlibabaHealthVaccinOrderUpdateAPIRequest struct {
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

// NewAlibabaHealthVaccinOrderUpdateRequest 初始化AlibabaHealthVaccinOrderUpdateAPIRequest对象
func NewAlibabaHealthVaccinOrderUpdateRequest() *AlibabaHealthVaccinOrderUpdateAPIRequest {
	return &AlibabaHealthVaccinOrderUpdateAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHealthVaccinOrderUpdateAPIRequest) Reset() {
	r._actionType = ""
	r._orderId = ""
	r._isvNotInoculateReason = ""
	r._actionTime = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinOrderUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.order.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinOrderUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHealthVaccinOrderUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActionType is ActionType Setter
// 操作类型
func (r *AlibabaHealthVaccinOrderUpdateAPIRequest) SetActionType(_actionType string) error {
	r._actionType = _actionType
	r.Set("action_type", _actionType)
	return nil
}

// GetActionType ActionType Getter
func (r AlibabaHealthVaccinOrderUpdateAPIRequest) GetActionType() string {
	return r._actionType
}

// SetOrderId is OrderId Setter
// 预约单id
func (r *AlibabaHealthVaccinOrderUpdateAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHealthVaccinOrderUpdateAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetIsvNotInoculateReason is IsvNotInoculateReason Setter
// 未接种原因
func (r *AlibabaHealthVaccinOrderUpdateAPIRequest) SetIsvNotInoculateReason(_isvNotInoculateReason string) error {
	r._isvNotInoculateReason = _isvNotInoculateReason
	r.Set("isv_not_inoculate_reason", _isvNotInoculateReason)
	return nil
}

// GetIsvNotInoculateReason IsvNotInoculateReason Getter
func (r AlibabaHealthVaccinOrderUpdateAPIRequest) GetIsvNotInoculateReason() string {
	return r._isvNotInoculateReason
}

// SetActionTime is ActionTime Setter
// 时间戳
func (r *AlibabaHealthVaccinOrderUpdateAPIRequest) SetActionTime(_actionTime int64) error {
	r._actionTime = _actionTime
	r.Set("action_time", _actionTime)
	return nil
}

// GetActionTime ActionTime Getter
func (r AlibabaHealthVaccinOrderUpdateAPIRequest) GetActionTime() int64 {
	return r._actionTime
}

var poolAlibabaHealthVaccinOrderUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHealthVaccinOrderUpdateRequest()
	},
}

// GetAlibabaHealthVaccinOrderUpdateRequest 从 sync.Pool 获取 AlibabaHealthVaccinOrderUpdateAPIRequest
func GetAlibabaHealthVaccinOrderUpdateAPIRequest() *AlibabaHealthVaccinOrderUpdateAPIRequest {
	return poolAlibabaHealthVaccinOrderUpdateAPIRequest.Get().(*AlibabaHealthVaccinOrderUpdateAPIRequest)
}

// ReleaseAlibabaHealthVaccinOrderUpdateAPIRequest 将 AlibabaHealthVaccinOrderUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaHealthVaccinOrderUpdateAPIRequest(v *AlibabaHealthVaccinOrderUpdateAPIRequest) {
	v.Reset()
	poolAlibabaHealthVaccinOrderUpdateAPIRequest.Put(v)
}
