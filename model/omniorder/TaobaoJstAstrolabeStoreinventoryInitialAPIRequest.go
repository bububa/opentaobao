package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstAstrolabeStoreinventoryInitialAPIRequest
后端商品库存初始化 API请求
taobao.jst.astrolabe.storeinventory.initial

初始化电商仓或门店库存，该接口一次可以初始化多个门店(或电商仓)的多个商品的多种类型库存。此接口只能使用一次，后续所有的库存变动都需走增量库存同步接口。 */
type TaobaoJstAstrolabeStoreinventoryInitialAPIRequest struct {
	model.Params
	// 操作时间
	_operationTime string
	// 门店列表
	_stores []Store
}

// New
