package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardReassign 工单改派门店
// tmall.servicecenter.workcard.reassign
//
// 工单改派门店
func TmallServicecenterWorkcardReassign(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardReassignAPIRequest, resp *tmallservice.TmallServicecenterWorkcardReassignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
