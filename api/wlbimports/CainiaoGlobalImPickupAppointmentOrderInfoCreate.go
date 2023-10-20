package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupAppointmentOrderInfoCreate 首公里揽收-预约单创建
// cainiao.global.im.pickup.appointment.order.info.create
//
// 预约单创建
func CainiaoGlobalImPickupAppointmentOrderInfoCreate(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIRequest, session string) (*wlbimports.CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIResponse, error) {
	var resp wlbimports.CainiaoGlobalImPickupAppointmentOrderInfoCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
