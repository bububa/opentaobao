package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenChangeticket 大麦换验平台-第三方对外开放-票单接口changeTicket
// alibaba.damai.mev.open.changeticket
//
// 开放接口 换票
func AlibabaDamaiMevOpenChangeticket(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenChangeticketAPIRequest, resp *damai.AlibabaDamaiMevOpenChangeticketAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
