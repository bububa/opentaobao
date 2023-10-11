package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaTmallSparepartsDetailsCreate 天猫蚁巢同步工单申请备件明细
// alibaba.tmall.spareparts.details.create
//
// 天猫蚁巢同步工单申请备件明细
func AlibabaTmallSparepartsDetailsCreate(clt *core.SDKClient, req *tmallsc.AlibabaTmallSparepartsDetailsCreateAPIRequest, session string) (*tmallsc.AlibabaTmallSparepartsDetailsCreateAPIResponse, error) {
	var resp tmallsc.AlibabaTmallSparepartsDetailsCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
