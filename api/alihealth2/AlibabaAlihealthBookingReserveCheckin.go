package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthBookingReserveCheckin 确认到店
// alibaba.alihealth.booking.reserve.checkin
//
// 消费医疗统一预约平台，ISV 确认到店
func AlibabaAlihealthBookingReserveCheckin(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBookingReserveCheckinAPIRequest, resp *alihealth2.AlibabaAlihealthBookingReserveCheckinAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
