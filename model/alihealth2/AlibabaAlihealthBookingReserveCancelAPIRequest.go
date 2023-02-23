package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBookingReserveCancelAPIRequest 取消预约 API请求
// alibaba.alihealth.booking.reserve.cancel
//
// 消费医疗统一预约平台，ISV取消预约
type AlibabaAlihealthBookingReserveCancelAPIRequest struct {
	model.Params
	// cancel
	_cancel *IsvReserveRequest
}

// NewAlibabaAlihealthBookingReserveCancelRequest 初始化AlibabaAlihealthBookingReserveCancelAPIRequest对象
func NewAlibabaAlihealthBookingReserveCancelRequest() *AlibabaAlihealthBookingReserveCancelAPIRequest {
	return &AlibabaAlihealthBookingReserveCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.booking.reserve.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthBookingReserveCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancel is Cancel Setter
// cancel
func (r *AlibabaAlihealthBookingReserveCancelAPIRequest) SetCancel(_cancel *IsvReserveRequest) error {
	r._cancel = _cancel
	r.Set("cancel", _cancel)
	return nil
}

// GetCancel Cancel Getter
func (r AlibabaAlihealthBookingReserveCancelAPIRequest) GetCancel() *IsvReserveRequest {
	return r._cancel
}
