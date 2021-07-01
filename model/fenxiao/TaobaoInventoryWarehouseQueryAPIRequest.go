package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryWarehouseQueryAPIRequest
分页查询商家仓信息 API请求
taobao.inventory.warehouse.query

分页查询商家仓信息 */
type TaobaoInventoryWarehouseQueryAPIRequest struct {
	model.Params
	// 页码
	_pageNo int64
	// 页大小
	_pageSize int64
}

// New
