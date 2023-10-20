package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitripaxintransfundquerybyorder 通过外部订单ID查询所有资金单
// taobao.alitrip.axin.trans.fund.query.by.order
//
// 阿信供销平台-通过外部订单ID查询所有资金单
func Taobaoalitripaxintransfundquerybyorder(clt *core.SDKClient, req *axintrade.TaobaoalitripaxintransfundquerybyorderAPIRequest, session string) (*axintrade.TaobaoalitripaxintransfundquerybyorderAPIResponse, error) {
	var resp axintrade.TaobaoalitripaxintransfundquerybyorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
