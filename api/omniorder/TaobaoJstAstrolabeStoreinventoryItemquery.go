package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoJstAstrolabeStoreinventoryItemquery 库存查询接口
// taobao.jst.astrolabe.storeinventory.itemquery
//
// 查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
func TaobaoJstAstrolabeStoreinventoryItemquery(clt *core.SDKClient, req *omniorder.TaobaoJstAstrolabeStoreinventoryItemqueryAPIRequest, resp *omniorder.TaobaoJstAstrolabeStoreinventoryItemqueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
