package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopenresetticket 大麦换验平台-第三方对外开放-票单接口resetTicket
// alibaba.damai.mev.open.resetticket
//
// 开放接口重打票
func Alibabadamaimevopenresetticket(clt *core.SDKClient, req *damai.AlibabadamaimevopenresetticketAPIRequest, session string) (*damai.AlibabadamaimevopenresetticketAPIResponse, error) {
	var resp damai.AlibabadamaimevopenresetticketAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
