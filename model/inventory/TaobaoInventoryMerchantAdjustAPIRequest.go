package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoInventoryMerchantAdjustAPIRequest
货品库存商家端调整 API请求
taobao.inventory.merchant.adjust

货品库存商家端调整 ，入库，出库，盘点 */
type TaobaoInventoryMerchantAdjustAPIRequest struct {
	model.Params
	// 调整库存对象
	_inventoryCheck *InventoryCheckDto
}

// New
