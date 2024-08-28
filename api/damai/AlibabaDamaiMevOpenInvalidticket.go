package damai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenInvalidticket 大麦换验平台-第三方对外开放-票单接口invalidTicket
// alibaba.damai.mev.open.invalidticket
//
// 开放接口 使票无效
func AlibabaDamaiMevOpenInvalidticket(ctx context.Context, clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenInvalidticketAPIRequest, resp *damai.AlibabaDamaiMevOpenInvalidticketAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
