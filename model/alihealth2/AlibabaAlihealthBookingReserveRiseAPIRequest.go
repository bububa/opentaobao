package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthBookingReserveRiseAPIRequest
ISV 新增/修改复诊预约信息 API请求
alibaba.alihealth.booking.reserve.rise

ISV 新增/修改复诊预约信息 */
type AlibabaAlihealthBookingReserveRiseAPIRequest struct {
	model.Params
	// 参数
	_riseRequest *IsvRiseReserveRequest
}

// New
