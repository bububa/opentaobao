package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthBookingReserveConfirmAPIRequest
确认预约 API请求
alibaba.alihealth.booking.reserve.confirm

确认预约 */
type AlibabaAlihealthBookingReserveConfirmAPIRequest struct {
	model.Params
	// 参数
	_confirm *IsvReserveRequest
}

// New
