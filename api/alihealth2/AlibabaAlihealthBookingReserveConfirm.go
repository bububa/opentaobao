package alihealth2

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthBookingReserveConfirm 确认预约
// alibaba.alihealth.booking.reserve.confirm
//
// 确认预约
func AlibabaAlihealthBookingReserveConfirm(ctx context.Context, clt *core.SDKClient, req *alihealth2.AlibabaAlihealthBookingReserveConfirmAPIRequest, resp *alihealth2.AlibabaAlihealthBookingReserveConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
