package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Cainiaoglobalimpickupappointmentordercancel 首公里揽收-取消预约单
// cainiao.global.im.pickup.appointment.order.cancel
//
// 首公里揽收-取消预约单创建
func Cainiaoglobalimpickupappointmentordercancel(clt *core.SDKClient, req *wlbimports.CainiaoglobalimpickupappointmentordercancelAPIRequest, session string) (*wlbimports.CainiaoglobalimpickupappointmentordercancelAPIResponse, error) {
	var resp wlbimports.CainiaoglobalimpickupappointmentordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
