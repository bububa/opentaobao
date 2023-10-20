package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Taobaoservindustryfinancegeexorderloan 即科放款信息回调api
// taobao.servindustry.finance.geex.order.loan
//
// 即科放款信息api
func Taobaoservindustryfinancegeexorderloan(clt *core.SDKClient, req *trade.TaobaoservindustryfinancegeexorderloanAPIRequest, session string) (*trade.TaobaoservindustryfinancegeexorderloanAPIResponse, error) {
	var resp trade.TaobaoservindustryfinancegeexorderloanAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
