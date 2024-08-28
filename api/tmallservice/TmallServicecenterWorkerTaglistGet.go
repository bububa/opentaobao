package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkerTaglistGet 获取工人标签
// tmall.servicecenter.worker.taglist.get
//
// 服务商获取对应工人的标签
func TmallServicecenterWorkerTaglistGet(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkerTaglistGetAPIRequest, resp *tmallservice.TmallServicecenterWorkerTaglistGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
