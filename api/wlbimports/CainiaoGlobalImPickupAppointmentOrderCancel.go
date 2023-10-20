package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupAppointmentOrderCancel 首公里揽收-取消预约单
// cainiao.global.im.pickup.appointment.order.cancel
//
// 首公里揽收-取消预约单创建
func CainiaoGlobalImPickupAppointmentOrderCancel(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupAppointmentOrderCancelAPIRequest, session string) (*wlbimports.CainiaoGlobalImPickupAppointmentOrderCancelAPIResponse, error) {
	var resp wlbimports.CainiaoGlobalImPickupAppointmentOrderCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
