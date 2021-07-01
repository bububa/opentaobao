package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest
库存增量更新接口 API请求
taobao.jst.astrolabe.storeinventory.itemupdate

ERP调用该接口，增量更新门店或电商仓库存，该接口一次可以同时增量更新多个门店的多个商品的非确定性库存。 */
type TaobaoJstAstrolabeStoreinventoryItemupdateAPIRequest struct {
	model.Params
	// 门店列表
	_stores []Store
	// 操作时间
	_operationTime string
}

// New
