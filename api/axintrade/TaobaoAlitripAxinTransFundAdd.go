package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitripaxintransfundadd 创建资金单接口
// taobao.alitrip.axin.trans.fund.add
//
// 创建资金单
func Taobaoalitripaxintransfundadd(clt *core.SDKClient, req *axintrade.TaobaoalitripaxintransfundaddAPIRequest, session string) (*axintrade.TaobaoalitripaxintransfundaddAPIResponse, error) {
	var resp axintrade.TaobaoalitripaxintransfundaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
