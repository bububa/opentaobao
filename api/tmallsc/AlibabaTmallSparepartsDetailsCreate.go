package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaTmallSparepartsDetailsCreate 天猫蚁巢同步工单申请备件明细
// alibaba.tmall.spareparts.details.create
//
// 天猫蚁巢同步工单申请备件明细
func AlibabaTmallSparepartsDetailsCreate(ctx context.Context, clt *core.SDKClient, req *tmallsc.AlibabaTmallSparepartsDetailsCreateAPIRequest, resp *tmallsc.AlibabaTmallSparepartsDetailsCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
