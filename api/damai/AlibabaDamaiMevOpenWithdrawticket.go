package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenWithdrawticket 大麦换验平台-第三方对外开放-票单接口withdrawTicket
// alibaba.damai.mev.open.withdrawticket
//
// 开放接口退票
func AlibabaDamaiMevOpenWithdrawticket(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenWithdrawticketAPIRequest, resp *damai.AlibabaDamaiMevOpenWithdrawticketAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
