package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Cainiaoglobalimpickupappointmentorderinfocreate 首公里揽收-预约单创建
// cainiao.global.im.pickup.appointment.order.info.create
//
// 预约单创建
func Cainiaoglobalimpickupappointmentorderinfocreate(clt *core.SDKClient, req *wlbimports.CainiaoglobalimpickupappointmentorderinfocreateAPIRequest, session string) (*wlbimports.CainiaoglobalimpickupappointmentorderinfocreateAPIResponse, error) {
	var resp wlbimports.CainiaoglobalimpickupappointmentorderinfocreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
