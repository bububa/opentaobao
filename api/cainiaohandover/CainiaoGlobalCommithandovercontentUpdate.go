package cainiaohandover

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalCommithandovercontentUpdate 修改已经提交的交接单
// cainiao.global.commithandovercontent.update
//
// 修改已经提交的交接单
func CainiaoGlobalCommithandovercontentUpdate(ctx context.Context, clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalCommithandovercontentUpdateAPIRequest, resp *cainiaohandover.CainiaoGlobalCommithandovercontentUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
