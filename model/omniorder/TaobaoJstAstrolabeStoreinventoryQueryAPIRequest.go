package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstAstrolabeStoreinventoryQueryAPIRequest
后端商品库存查询接口 API请求
taobao.jst.astrolabe.storeinventory.query

查询门店或电商仓库存，该接口一次可以同时查询多个门店或电商仓的多个商品的多种类型的库存 */
type TaobaoJstAstrolabeStoreinventoryQueryAPIRequest struct {
	model.Params
	// 门店
	_stores []Store
}

// New
