package damai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/damai"
)

// Alibabadamaimevopenunlockticket 大麦换验平台-第三方对外开放-票单接口unlockTicket
// alibaba.damai.mev.open.unlockticket
//
// 开放接口 解锁票单
func Alibabadamaimevopenunlockticket(clt *core.SDKClient, req *damai.AlibabadamaimevopenunlockticketAPIRequest, session string) (*damai.AlibabadamaimevopenunlockticketAPIResponse, error) {
	var resp damai.AlibabadamaimevopenunlockticketAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
