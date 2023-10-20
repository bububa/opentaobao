package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthBookingReserveRise ISV 新增/修改复诊预约信息
// alibaba.alihealth.booking.reserve.rise
//
// ISV 新增/修改复诊预约信息
func AlibabaAlihealthBookingReserveRise(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBookingReserveRiseAPIRequest, resp *alihealth2.AlibabaAlihealthBookingReserveRiseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
