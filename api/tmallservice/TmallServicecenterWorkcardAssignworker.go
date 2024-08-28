package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardAssignworker 服务商分派工人
// tmall.servicecenter.workcard.assignworker
//
// 服务商调用该接口分派工人给具体的工单
func TmallServicecenterWorkcardAssignworker(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardAssignworkerAPIRequest, resp *tmallservice.TmallServicecenterWorkcardAssignworkerAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
