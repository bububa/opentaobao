package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoJstAstrolabeStoreinventoryQuery 后端商品库存查询接口
// taobao.jst.astrolabe.storeinventory.query
//
// 查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存
func TaobaoJstAstrolabeStoreinventoryQuery(clt *core.SDKClient, req *omniorder.TaobaoJstAstrolabeStoreinventoryQueryAPIRequest, resp *omniorder.TaobaoJstAstrolabeStoreinventoryQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
