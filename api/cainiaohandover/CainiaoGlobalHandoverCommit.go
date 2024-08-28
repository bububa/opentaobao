package cainiaohandover

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// CainiaoGlobalHandoverCommit 提交发布交接单
// cainiao.global.handover.commit
//
// 提供给ISV通过该接口提交发布交接单
func CainiaoGlobalHandoverCommit(ctx context.Context, clt *core.SDKClient, req *cainiaohandover.CainiaoGlobalHandoverCommitAPIRequest, resp *cainiaohandover.CainiaoGlobalHandoverCommitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
