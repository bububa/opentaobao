package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstAstrolabeStoreinventoryUpdateAPIRequest
后端商品库存增量更新接口 API请求
taobao.jst.astrolabe.storeinventory.update

增量更新门店或电商仓库存，该接口一次可以同时增量更新多个门店的多个商品的非确定性库存 */
type TaobaoJstAstrolabeStoreinventoryUpdateAPIRequest struct {
	model.Params
	// 操作时间
	_operationTime string
	// 门店列表
	_stores []Store
}

// New
