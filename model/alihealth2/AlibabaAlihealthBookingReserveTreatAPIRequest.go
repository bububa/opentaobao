package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBookingReserveTreatAPIRequest 确认就诊 API请求
// alibaba.alihealth.booking.reserve.treat
//
// 统一预约平台，ISV确认就诊
type AlibabaAlihealthBookingReserveTreatAPIRequest struct {
	model.Params
	// treat
	_treat *IsvReserveRequest
}

// NewAlibabaAlihealthBookingReserveTreatRequest 初始化AlibabaAlihealthBookingReserveTreatAPIRequest对象
func NewAlibabaAlihealthBookingReserveTreatRequest() *AlibabaAlihealthBookingReserveTreatAPIRequest {
	return &AlibabaAlihealthBookingReserveTreatAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthBookingReserveTreatAPIRequest) Reset() {
	r._treat = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveTreatAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.booking.reserve.treat"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveTreatAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthBookingReserveTreatAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTreat is Treat Setter
// treat
func (r *AlibabaAlihealthBookingReserveTreatAPIRequest) SetTreat(_treat *IsvReserveRequest) error {
	r._treat = _treat
	r.Set("treat", _treat)
	return nil
}

// GetTreat Treat Getter
func (r AlibabaAlihealthBookingReserveTreatAPIRequest) GetTreat() *IsvReserveRequest {
	return r._treat
}

var poolAlibabaAlihealthBookingReserveTreatAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthBookingReserveTreatRequest()
	},
}

// GetAlibabaAlihealthBookingReserveTreatRequest 从 sync.Pool 获取 AlibabaAlihealthBookingReserveTreatAPIRequest
func GetAlibabaAlihealthBookingReserveTreatAPIRequest() *AlibabaAlihealthBookingReserveTreatAPIRequest {
	return poolAlibabaAlihealthBookingReserveTreatAPIRequest.Get().(*AlibabaAlihealthBookingReserveTreatAPIRequest)
}

// ReleaseAlibabaAlihealthBookingReserveTreatAPIRequest 将 AlibabaAlihealthBookingReserveTreatAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthBookingReserveTreatAPIRequest(v *AlibabaAlihealthBookingReserveTreatAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthBookingReserveTreatAPIRequest.Put(v)
}
