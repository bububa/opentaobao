package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkerQuerypage 查询工人列表
// tmall.servicecenter.worker.querypage
//
// 服务商查询工人列表
func TmallServicecenterWorkerQuerypage(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkerQuerypageAPIRequest, resp *tmallservice.TmallServicecenterWorkerQuerypageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
