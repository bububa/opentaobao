package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeStoreinventoryItemadjustAPIRequest 库存占用调整接口 API请求
// taobao.jst.astrolabe.storeinventory.itemadjust
//
// 当第三方系统出现分单结果和天猫货品中心分单结果不一致时，需要调用此接口同步分单消息给天猫货品中心，调整之前占用的门店/电商仓库存。
type TaobaoJstAstrolabeStoreinventoryItemadjustAPIRequest struct {
	model.Params
	// 操作时间
	_operationTime string
	// 库存调整信息
	_inventoryAdjustInfo *InventoryAdjustInfo
}

// NewTaobaoJstAstrolabeStoreinventoryItemadjustRequest 初始化TaobaoJstAstrolabeStoreinventoryItemadjustAPIRequest对象
func NewTaobaoJstAstrolabeStoreinventoryItemadjustRequest() *TaobaoJstAstrolabeStoreinventoryItemadjustAPIRequest {
	return &TaobaoJstAstrolabeStoreinventoryItemadjustAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeStoreinventoryItemadjustAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.itemadjust"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeStoreinventoryItemadjustAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOperationTime is OperationTime Setter
// 操作时间
func (r *TaobaoJstAstrolabeStoreinventoryItemadjustAPIRequest) SetOperationTime(_operationTime string) error {
	r._operationTime = _operationTime
	r.Set("operation_time", _operationTime)
	return nil
}

// GetOperationTime OperationTime Getter
func (r TaobaoJstAstrolabeStoreinventoryItemadjustAPIRequest) GetOperationTime() string {
	return r._operationTime
}

// SetInventoryAdjustInfo is InventoryAdjustInfo Setter
// 库存调整信息
func (r *TaobaoJstAstrolabeStoreinventoryItemadjustAPIRequest) SetInventoryAdjustInfo(_inventoryAdjustInfo *InventoryAdjustInfo) error {
	r._inventoryAdjustInfo = _inventoryAdjustInfo
	r.Set("inventory_adjust_info", _inventoryAdjustInfo)
	return nil
}

// GetInventoryAdjustInfo InventoryAdjustInfo Getter
func (r TaobaoJstAstrolabeStoreinventoryItemadjustAPIRequest) GetInventoryAdjustInfo() *InventoryAdjustInfo {
	return r._inventoryAdjustInfo
}
