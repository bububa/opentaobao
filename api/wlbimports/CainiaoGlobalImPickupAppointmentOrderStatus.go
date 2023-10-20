package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupAppointmentOrderStatus 预约单状态查询
// cainiao.global.im.pickup.appointment.order.status
//
// 预约单状态查询
func CainiaoGlobalImPickupAppointmentOrderStatus(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest, resp *wlbimports.CainiaoGlobalImPickupAppointmentOrderStatusAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
