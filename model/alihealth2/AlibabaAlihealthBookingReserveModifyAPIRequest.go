package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthBookingReserveModifyAPIRequest
修改预约 API请求
alibaba.alihealth.booking.reserve.modify

消费医疗统一预约平台，取消预约 */
type AlibabaAlihealthBookingReserveModifyAPIRequest struct {
	model.Params
	// modify
	_modify *IsvReserveRequest
}

// New
