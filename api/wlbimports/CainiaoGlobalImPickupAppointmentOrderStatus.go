package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupAppointmentOrderStatus 预约单状态查询
// cainiao.global.im.pickup.appointment.order.status
//
// 预约单状态查询
func CainiaoGlobalImPickupAppointmentOrderStatus(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupAppointmentOrderStatusAPIRequest, session string) (*wlbimports.CainiaoGlobalImPickupAppointmentOrderStatusAPIResponse, error) {
	var resp wlbimports.CainiaoGlobalImPickupAppointmentOrderStatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
