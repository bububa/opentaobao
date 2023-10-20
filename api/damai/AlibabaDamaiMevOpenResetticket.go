package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenResetticket 大麦换验平台-第三方对外开放-票单接口resetTicket
// alibaba.damai.mev.open.resetticket
//
// 开放接口重打票
func AlibabaDamaiMevOpenResetticket(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenResetticketAPIRequest, resp *damai.AlibabaDamaiMevOpenResetticketAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
