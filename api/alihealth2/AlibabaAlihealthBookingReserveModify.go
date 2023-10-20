package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthBookingReserveModify 修改预约
// alibaba.alihealth.booking.reserve.modify
//
// 消费医疗统一预约平台，取消预约
func AlibabaAlihealthBookingReserveModify(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBookingReserveModifyAPIRequest, resp *alihealth2.AlibabaAlihealthBookingReserveModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
