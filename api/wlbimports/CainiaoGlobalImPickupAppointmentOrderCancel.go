package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupAppointmentOrderCancel 首公里揽收-取消预约单
// cainiao.global.im.pickup.appointment.order.cancel
//
// 首公里揽收-取消预约单创建
func CainiaoGlobalImPickupAppointmentOrderCancel(ctx context.Context, clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest, resp *wlbimports.CainiaoGlobalImPickupAppointmentOrderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
