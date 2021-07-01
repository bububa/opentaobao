package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryWarehouseManageAPIRequest
创建商家仓或者更新商家仓信息 API请求
taobao.inventory.warehouse.manage

创建商家仓或者更新商家仓信息 */
type TaobaoInventoryWarehouseManageAPIRequest struct {
	model.Params
	// 仓库信息
	_wareHouseDto *WareHouseDto
}

// New
