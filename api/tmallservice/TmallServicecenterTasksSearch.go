package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterTasksSearch 查询任务类工单信息
// tmall.servicecenter.tasks.search
//
// 查询任务类工单信息
func TmallServicecenterTasksSearch(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterTasksSearchAPIRequest, resp *tmallservice.TmallServicecenterTasksSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
