package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthBookingReserveTreat 确认就诊
// alibaba.alihealth.booking.reserve.treat
//
// 统一预约平台，ISV确认就诊
func AlibabaAlihealthBookingReserveTreat(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBookingReserveTreatAPIRequest, resp *alihealth2.AlibabaAlihealthBookingReserveTreatAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
