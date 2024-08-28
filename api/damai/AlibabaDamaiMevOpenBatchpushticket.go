package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenBatchpushticket 大麦换验平台-第三方对外开放-票单接口batchPushTicket
// alibaba.damai.mev.open.batchpushticket
//
// 批量推送票单
func AlibabaDamaiMevOpenBatchpushticket(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenBatchpushticketAPIRequest, resp *damai.AlibabaDamaiMevOpenBatchpushticketAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
