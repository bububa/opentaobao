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

// NewTaobaoJstAstrolabeStoreinventoryAdjustRequest 初始化TaobaoJstAstrolabeStoreinventoryAdjustAPIRequest对象
func NewTaobaoJstAstrolabeStoreinventoryAdjustRequest() *TaobaoJstAstrolabeStoreinventoryAdjustAPIRequest {
	return &TaobaoJstAstrolabeStoreinventoryAdjustAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryAdjustAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.adjust"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryAdjustAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OperationTime Setter
// 操作时间
func (r *TaobaoJstAstrolabeStoreinventoryAdjustAPIRequest) SetOperationTime(_operationTime string) error {
	r._operationTime = _operationTime
	r.Set("operation_time", _operationTime)
	return nil
}

// Get OperationTime Getter
func (r TaobaoJstAstrolabeStoreinventoryAdjustAPIRequest) GetOperationTime() string {
	return r._operationTime
}

// Set is InventoryAdjustInfo Setter
// 库存调整信息
func (r *TaobaoJstAstrolabeStoreinventoryAdjustAPIRequest) SetInventoryAdjustInfo(_inventoryAdjustInfo *InventoryAdjustInfo) error {
	r._inventoryAdjustInfo = _inventoryAdjustInfo
	r.Set("inventory_adjust_info", _inventoryAdjustInfo)
	return nil
}

// Get InventoryAdjustInfo Getter
func (r TaobaoJstAstrolabeStoreinventoryAdjustAPIRequest) GetInventoryAdjustInfo() *InventoryAdjustInfo {
	return r._inventoryAdjustInfo
}
