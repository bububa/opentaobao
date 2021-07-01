package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstAstrolabeStoreinventoryAdjustAPIRequest
后端商品库存占用调整接口 API请求
taobao.jst.astrolabe.storeinventory.adjust

当第三方系统出现分单结果和天猫货品中心分单结果不一致时，需要调用此接口同步分单消息给天猫货品中心，调整之前占用的门店/电商仓库存。 */
type TaobaoJstAstrolabeStoreinventoryAdjustAPIRequest struct {
	model.Params
	// 操作时间
	_operationTime string
	// 库存调整信息
	_inventoryAdjustInfo *InventoryAdjustInfo
}

// New
