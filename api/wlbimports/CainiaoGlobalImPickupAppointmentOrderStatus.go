package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Cainiaoglobalimpickupappointmentorderstatus 预约单状态查询
// cainiao.global.im.pickup.appointment.order.status
//
// 预约单状态查询
func Cainiaoglobalimpickupappointmentorderstatus(clt *core.SDKClient, req *wlbimports.CainiaoglobalimpickupappointmentorderstatusAPIRequest, session string) (*wlbimports.CainiaoglobalimpickupappointmentorderstatusAPIResponse, error) {
	var resp wlbimports.CainiaoglobalimpickupappointmentorderstatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
