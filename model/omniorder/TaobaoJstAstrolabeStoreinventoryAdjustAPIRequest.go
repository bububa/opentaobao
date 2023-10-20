package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstastrolabestoreinventoryadjustAPIRequest 后端商品库存占用调整接口 API请求
// taobao.jst.astrolabe.storeinventory.adjust
//
// 当第三方系统出现分单结果和天猫货品中心分单结果不一致时，需要调用此接口同步分单消息给天猫货品中心，调整之前占用的门店/电商仓库存。
type TaobaojstastrolabestoreinventoryadjustAPIRequest struct {
	model.Params
	// 操作时间
	_operationTime string
	// 库存调整信息
	_inventoryAdjustInfo *InventoryAdjustInfo
}

// NewTaobaojstastrolabestoreinventoryadjustRequest 初始化TaobaojstastrolabestoreinventoryadjustAPIRequest对象
func NewTaobaojstastrolabestoreinventoryadjustRequest() *TaobaojstastrolabestoreinventoryadjustAPIRequest {
	return &TaobaojstastrolabestoreinventoryadjustAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstastrolabestoreinventoryadjustAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.adjust"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstastrolabestoreinventoryadjustAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstastrolabestoreinventoryadjustAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOperationTime is OperationTime Setter
// 操作时间
func (r *TaobaojstastrolabestoreinventoryadjustAPIRequest) SetOperationTime(_operationTime string) error {
	r._operationTime = _operationTime
	r.Set("operation_time", _operationTime)
	return nil
}

// GetOperationTime OperationTime Getter
func (r TaobaojstastrolabestoreinventoryadjustAPIRequest) GetOperationTime() string {
	return r._operationTime
}

// SetInventoryAdjustInfo is InventoryAdjustInfo Setter
// 库存调整信息
func (r *TaobaojstastrolabestoreinventoryadjustAPIRequest) SetInventoryAdjustInfo(_inventoryAdjustInfo *InventoryAdjustInfo) error {
	r._inventoryAdjustInfo = _inventoryAdjustInfo
	r.Set("inventory_adjust_info", _inventoryAdjustInfo)
	return nil
}

// GetInventoryAdjustInfo InventoryAdjustInfo Getter
func (r TaobaojstastrolabestoreinventoryadjustAPIRequest) GetInventoryAdjustInfo() *InventoryAdjustInfo {
	return r._inventoryAdjustInfo
}
