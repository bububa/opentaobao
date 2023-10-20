package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopeninvalidticket 大麦换验平台-第三方对外开放-票单接口invalidTicket
// alibaba.damai.mev.open.invalidticket
//
// 开放接口 使票无效
func Alibabadamaimevopeninvalidticket(clt *core.SDKClient, req *damai.AlibabadamaimevopeninvalidticketAPIRequest, session string) (*damai.AlibabadamaimevopeninvalidticketAPIResponse, error) {
	var resp damai.AlibabadamaimevopeninvalidticketAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
