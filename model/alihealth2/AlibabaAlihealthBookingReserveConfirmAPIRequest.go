package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBookingReserveConfirmAPIRequest 确认预约 API请求
// alibaba.alihealth.booking.reserve.confirm
//
// 确认预约
type AlibabaAlihealthBookingReserveConfirmAPIRequest struct {
	model.Params
	// 参数
	_confirm *IsvReserveRequest
}

// NewAlibabaAlihealthBookingReserveConfirmRequest 初始化AlibabaAlihealthBookingReserveConfirmAPIRequest对象
func NewAlibabaAlihealthBookingReserveConfirmRequest() *AlibabaAlihealthBookingReserveConfirmAPIRequest {
	return &AlibabaAlihealthBookingReserveConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthBookingReserveConfirmAPIRequest) Reset() {
	r._confirm = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.booking.reserve.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthBookingReserveConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConfirm is Confirm Setter
// 参数
func (r *AlibabaAlihealthBookingReserveConfirmAPIRequest) SetConfirm(_confirm *IsvReserveRequest) error {
	r._confirm = _confirm
	r.Set("confirm", _confirm)
	return nil
}

// GetConfirm Confirm Getter
func (r AlibabaAlihealthBookingReserveConfirmAPIRequest) GetConfirm() *IsvReserveRequest {
	return r._confirm
}

var poolAlibabaAlihealthBookingReserveConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthBookingReserveConfirmRequest()
	},
}

// GetAlibabaAlihealthBookingReserveConfirmRequest 从 sync.Pool 获取 AlibabaAlihealthBookingReserveConfirmAPIRequest
func GetAlibabaAlihealthBookingReserveConfirmAPIRequest() *AlibabaAlihealthBookingReserveConfirmAPIRequest {
	return poolAlibabaAlihealthBookingReserveConfirmAPIRequest.Get().(*AlibabaAlihealthBookingReserveConfirmAPIRequest)
}

// ReleaseAlibabaAlihealthBookingReserveConfirmAPIRequest 将 AlibabaAlihealthBookingReserveConfirmAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthBookingReserveConfirmAPIRequest(v *AlibabaAlihealthBookingReserveConfirmAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthBookingReserveConfirmAPIRequest.Put(v)
}
