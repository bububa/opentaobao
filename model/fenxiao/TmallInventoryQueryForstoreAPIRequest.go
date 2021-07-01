package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallInventoryQueryForstoreAPIRequest
查询后端商品仓库库存 API请求
tmall.inventory.query.forstore

商家查询后端商品仓库库存 */
type TmallInventoryQueryForstoreAPIRequest struct {
	model.Params
	// 查询列表
	_paramList []InventoryQueryForStoreRequest
}

// New
