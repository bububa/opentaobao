package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// AlibabaDamaiMevOpenInvalidticket 大麦换验平台-第三方对外开放-票单接口invalidTicket
// alibaba.damai.mev.open.invalidticket
//
// 开放接口 使票无效
func AlibabaDamaiMevOpenInvalidticket(clt *core.SDKClient, req *damai.AlibabaDamaiMevOpenInvalidticketAPIRequest, session string) (*damai.AlibabaDamaiMevOpenInvalidticketAPIResponse, error) {
	var resp damai.AlibabaDamaiMevOpenInvalidticketAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
