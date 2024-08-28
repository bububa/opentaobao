package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardComplete 工单完结
// tmall.servicecenter.workcard.complete
//
// 工单完结
func TmallServicecenterWorkcardComplete(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardCompleteAPIRequest, resp *tmallservice.TmallServicecenterWorkcardCompleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
