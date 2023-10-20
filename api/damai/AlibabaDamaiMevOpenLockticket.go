package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenLockticket 大麦换验平台-第三方对外开放-票单接口lockTicket
// alibaba.damai.mev.open.lockticket
//
// 开放接口 冻结票单
func AlibabaDamaiMevOpenLockticket(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenLockticketAPIRequest, resp *damai.AlibabaDamaiMevOpenLockticketAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
