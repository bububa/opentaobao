package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractOnecodeIssue onecode代码通用鉴权
// alibaba.interact.onecode.issue
//
// 手淘开放鉴权接口，仅用于tida接口鉴权，无输入输出。
func AlibabaInteractOnecodeIssue(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractOnecodeIssueAPIRequest, resp *interact.AlibabaInteractOnecodeIssueAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
