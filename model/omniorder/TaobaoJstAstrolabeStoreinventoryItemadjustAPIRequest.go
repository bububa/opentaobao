package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstastrolabestoreinventoryitemadjustAPIRequest 库存占用调整接口 API请求
// taobao.jst.astrolabe.storeinventory.itemadjust
//
// 当第三方系统出现分单结果和天猫货品中心分单结果不一致时，需要调用此接口同步分单消息给天猫货品中心，调整之前占用的门店/电商仓库存。
type TaobaojstastrolabestoreinventoryitemadjustAPIRequest struct {
	model.Params
	// 操作时间
	_operationTime string
	// 库存调整信息
	_inventoryAdjustInfo *InventoryAdjustInfo
}

// NewTaobaojstastrolabestoreinventoryitemadjustRequest 初始化TaobaojstastrolabestoreinventoryitemadjustAPIRequest对象
func NewTaobaojstastrolabestoreinventoryitemadjustRequest() *TaobaojstastrolabestoreinventoryitemadjustAPIRequest {
	return &TaobaojstastrolabestoreinventoryitemadjustAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaojstastrolabestoreinventoryitemadjustAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.storeinventory.itemadjust"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaojstastrolabestoreinventoryitemadjustAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaojstastrolabestoreinventoryitemadjustAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOperationTime is OperationTime Setter
// 操作时间
func (r *TaobaojstastrolabestoreinventoryitemadjustAPIRequest) SetOperationTime(_operationTime string) error {
	r._operationTime = _operationTime
	r.Set("operation_time", _operationTime)
	return nil
}

// GetOperationTime OperationTime Getter
func (r TaobaojstastrolabestoreinventoryitemadjustAPIRequest) GetOperationTime() string {
	return r._operationTime
}

// SetInventoryAdjustInfo is InventoryAdjustInfo Setter
// 库存调整信息
func (r *TaobaojstastrolabestoreinventoryitemadjustAPIRequest) SetInventoryAdjustInfo(_inventoryAdjustInfo *InventoryAdjustInfo) error {
	r._inventoryAdjustInfo = _inventoryAdjustInfo
	r.Set("inventory_adjust_info", _inventoryAdjustInfo)
	return nil
}

// GetInventoryAdjustInfo InventoryAdjustInfo Getter
func (r TaobaojstastrolabestoreinventoryitemadjustAPIRequest) GetInventoryAdjustInfo() *InventoryAdjustInfo {
	return r._inventoryAdjustInfo
}
