package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenUnlockticket 大麦换验平台-第三方对外开放-票单接口unlockTicket
// alibaba.damai.mev.open.unlockticket
//
// 开放接口 解锁票单
func AlibabaDamaiMevOpenUnlockticket(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenUnlockticketAPIRequest, resp *damai.AlibabaDamaiMevOpenUnlockticketAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
