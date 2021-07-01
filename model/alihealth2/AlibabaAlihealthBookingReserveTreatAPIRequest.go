package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthBookingReserveTreatAPIRequest
确认就诊 API请求
alibaba.alihealth.booking.reserve.treat

统一预约平台，ISV确认就诊 */
type AlibabaAlihealthBookingReserveTreatAPIRequest struct {
	model.Params
	// treat
	_treat *IsvReserveRequest
}

// New
