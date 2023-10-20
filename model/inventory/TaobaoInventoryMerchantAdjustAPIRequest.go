package inventory

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoinventorymerchantadjustAPIRequest 货品库存商家端调整 API请求
// taobao.inventory.merchant.adjust
//
// 货品库存商家端调整 ，入库，出库，盘点
type TaobaoinventorymerchantadjustAPIRequest struct {
	model.Params
	// 调整库存对象
	_inventoryCheck *InventoryCheckDto
}

// NewTaobaoinventorymerchantadjustRequest 初始化TaobaoinventorymerchantadjustAPIRequest对象
func NewTaobaoinventorymerchantadjustRequest() *TaobaoinventorymerchantadjustAPIRequest {
	return &TaobaoinventorymerchantadjustAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoinventorymerchantadjustAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.merchant.adjust"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoinventorymerchantadjustAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoinventorymerchantadjustAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInventoryCheck is InventoryCheck Setter
// 调整库存对象
func (r *TaobaoinventorymerchantadjustAPIRequest) SetInventoryCheck(_inventoryCheck *InventoryCheckDto) error {
	r._inventoryCheck = _inventoryCheck
	r.Set("inventory_check", _inventoryCheck)
	return nil
}

// GetInventoryCheck InventoryCheck Getter
func (r TaobaoinventorymerchantadjustAPIRequest) GetInventoryCheck() *InventoryCheckDto {
	return r._inventoryCheck
}
