package alihealth2

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveConfirmAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.booking.reserve.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Confirm Setter
// 参数
func (r *AlibabaAlihealthBookingReserveConfirmAPIRequest) SetConfirm(_confirm *IsvReserveRequest) error {
	r._confirm = _confirm
	r.Set("confirm", _confirm)
	return nil
}

// Get Confirm Getter
func (r AlibabaAlihealthBookingReserveConfirmAPIRequest) GetConfirm() *IsvReserveRequest {
	return r._confirm
}
