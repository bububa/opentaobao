package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupAppointmentOrderDifferenceDetail 预约单差异明细查询
// cainiao.global.im.pickup.appointment.order.difference.detail
//
// 预约单差异明细查询
func CainiaoGlobalImPickupAppointmentOrderDifferenceDetail(ctx context.Context, clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIRequest, resp *wlbimports.CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
