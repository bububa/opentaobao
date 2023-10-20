package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuBuyerTagGet 判断买家是否有某些标
// taobao.qianniu.buyer.tag.get
//
// 判断某个买家是否有某些标
func TaobaoQianniuBuyerTagGet(clt *core.SDKClient, req *qianniu.TaobaoQianniuBuyerTagGetAPIRequest, resp *qianniu.TaobaoQianniuBuyerTagGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
