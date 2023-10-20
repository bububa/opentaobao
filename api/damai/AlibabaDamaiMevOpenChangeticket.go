package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopenchangeticket 大麦换验平台-第三方对外开放-票单接口changeTicket
// alibaba.damai.mev.open.changeticket
//
// 开放接口 换票
func Alibabadamaimevopenchangeticket(clt *core.SDKClient, req *damai.AlibabadamaimevopenchangeticketAPIRequest, session string) (*damai.AlibabadamaimevopenchangeticketAPIResponse, error) {
	var resp damai.AlibabadamaimevopenchangeticketAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
