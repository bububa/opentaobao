package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBookingReserveCheckinAPIRequest 确认到店 API请求
// alibaba.alihealth.booking.reserve.checkin
//
// 消费医疗统一预约平台，ISV 确认到店
type AlibabaAlihealthBookingReserveCheckinAPIRequest struct {
	model.Params
	// check_in
	_checkIn *IsvReserveRequest
}

// NewAlibabaAlihealthBookingReserveCheckinRequest 初始化AlibabaAlihealthBookingReserveCheckinAPIRequest对象
func NewAlibabaAlihealthBookingReserveCheckinRequest() *AlibabaAlihealthBookingReserveCheckinAPIRequest {
	return &AlibabaAlihealthBookingReserveCheckinAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthBookingReserveCheckinAPIRequest) Reset() {
	r._checkIn = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthBookingReserveCheckinAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.booking.reserve.checkin"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthBookingReserveCheckinAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthBookingReserveCheckinAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCheckIn is CheckIn Setter
// check_in
func (r *AlibabaAlihealthBookingReserveCheckinAPIRequest) SetCheckIn(_checkIn *IsvReserveRequest) error {
	r._checkIn = _checkIn
	r.Set("check_in", _checkIn)
	return nil
}

// GetCheckIn CheckIn Getter
func (r AlibabaAlihealthBookingReserveCheckinAPIRequest) GetCheckIn() *IsvReserveRequest {
	return r._checkIn
}

var poolAlibabaAlihealthBookingReserveCheckinAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthBookingReserveCheckinRequest()
	},
}

// GetAlibabaAlihealthBookingReserveCheckinRequest 从 sync.Pool 获取 AlibabaAlihealthBookingReserveCheckinAPIRequest
func GetAlibabaAlihealthBookingReserveCheckinAPIRequest() *AlibabaAlihealthBookingReserveCheckinAPIRequest {
	return poolAlibabaAlihealthBookingReserveCheckinAPIRequest.Get().(*AlibabaAlihealthBookingReserveCheckinAPIRequest)
}

// ReleaseAlibabaAlihealthBookingReserveCheckinAPIRequest 将 AlibabaAlihealthBookingReserveCheckinAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthBookingReserveCheckinAPIRequest(v *AlibabaAlihealthBookingReserveCheckinAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthBookingReserveCheckinAPIRequest.Put(v)
}
