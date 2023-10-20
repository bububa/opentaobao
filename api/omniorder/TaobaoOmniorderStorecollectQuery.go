package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStorecollectQuery 全渠道门店自提根据核销码查询订单
// taobao.omniorder.storecollect.query
//
// 全渠道门店自提根据核销码查询订单
func TaobaoOmniorderStorecollectQuery(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStorecollectQueryAPIRequest, resp *omniorder.TaobaoOmniorderStorecollectQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
