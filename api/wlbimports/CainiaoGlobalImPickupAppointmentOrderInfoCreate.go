package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupAppointmentOrderInfoCreate 首公里揽收-预约单创建
// cainiao.global.im.pickup.appointment.order.info.create
//
// 预约单创建
func CainiaoGlobalImPickupAppointmentOrderInfoCreate(ctx context.Context, clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest, resp *wlbimports.CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
