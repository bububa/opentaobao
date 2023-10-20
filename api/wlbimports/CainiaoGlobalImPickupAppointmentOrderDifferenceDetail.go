package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupAppointmentOrderDifferenceDetail 预约单差异明细查询
// cainiao.global.im.pickup.appointment.order.difference.detail
//
// 预约单差异明细查询
func CainiaoGlobalImPickupAppointmentOrderDifferenceDetail(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIRequest, session string) (*wlbimports.CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse, error) {
	var resp wlbimports.CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
