package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaTmallSparepartsDetailsCreate 天猫蚁巢同步工单申请备件明细
// alibaba.tmall.spareparts.details.create
//
// 天猫蚁巢同步工单申请备件明细
func AlibabaTmallSparepartsDetailsCreate(clt *core.SDKClient, req *tmallsc.AlibabaTmallSparepartsDetailsCreateAPIRequest, resp *tmallsc.AlibabaTmallSparepartsDetailsCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
