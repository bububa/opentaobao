package rhino

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRhinoSupplychainClothingAdjustAPIRequest
同步成衣仓盘点数据 API请求
taobao.rhino.supplychain.clothing.adjust

同步成衣仓盘点数据 */
type TaobaoRhinoSupplychainClothingAdjustAPIRequest struct {
	model.Params
	// 库存盘点对象
	_param0 *MaterialInventoryAdjustDto
}

// New
