package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthBookingReserveCancelAPIRequest
取消预约 API请求
alibaba.alihealth.booking.reserve.cancel

消费医疗统一预约平台，ISV取消预约 */
type AlibabaAlihealthBookingReserveCancelAPIRequest struct {
	model.Params
	// cancel
	_cancel *IsvReserveRequest
}

// New
