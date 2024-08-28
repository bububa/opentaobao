package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// AlibabaMsfserviceWorkerQueryid 查询师傅workerid
// alibaba.msfservice.worker.queryid
//
// 查询师傅workerid
func AlibabaMsfserviceWorkerQueryid(ctx context.Context, clt *core.SDKClient, req *tmallsc.AlibabaMsfserviceWorkerQueryidAPIRequest, resp *tmallsc.AlibabaMsfserviceWorkerQueryidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
