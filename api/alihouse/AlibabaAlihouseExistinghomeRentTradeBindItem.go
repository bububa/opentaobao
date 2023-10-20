package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomerenttradebinditem 交易绑定商品
// alibaba.alihouse.existinghome.rent.trade.bind.item
//
// 交易绑定商品
func Alibabaalihouseexistinghomerenttradebinditem(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomerenttradebinditemAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomerenttradebinditemAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomerenttradebinditemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
