package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthBookingReserveCheckinAPIRequest
确认到店 API请求
alibaba.alihealth.booking.reserve.checkin

消费医疗统一预约平台，ISV 确认到店 */
type AlibabaAlihealthBookingReserveCheckinAPIRequest struct {
	model.Params
	// check_in
	_checkIn *IsvReserveRequest
}

// New
