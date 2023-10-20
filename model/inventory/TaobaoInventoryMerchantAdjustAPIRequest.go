package inventory

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryMerchantAdjustAPIRequest 货品库存商家端调整 API请求
// taobao.inventory.merchant.adjust
//
// 货品库存商家端调整 ，入库，出库，盘点
type TaobaoInventoryMerchantAdjustAPIRequest struct {
	model.Params
	// 调整库存对象
	_inventoryCheck *InventoryCheckDto
}

// NewTaobaoInventoryMerchantAdjustRequest 初始化TaobaoInventoryMerchantAdjustAPIRequest对象
func NewTaobaoInventoryMerchantAdjustRequest() *TaobaoInventoryMerchantAdjustAPIRequest {
	return &TaobaoInventoryMerchantAdjustAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoInventoryMerchantAdjustAPIRequest) Reset() {
	r._inventoryCheck = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryMerchantAdjustAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.merchant.adjust"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryMerchantAdjustAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoInventoryMerchantAdjustAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetInventoryCheck is InventoryCheck Setter
// 调整库存对象
func (r *TaobaoInventoryMerchantAdjustAPIRequest) SetInventoryCheck(_inventoryCheck *InventoryCheckDto) error {
	r._inventoryCheck = _inventoryCheck
	r.Set("inventory_check", _inventoryCheck)
	return nil
}

// GetInventoryCheck InventoryCheck Getter
func (r TaobaoInventoryMerchantAdjustAPIRequest) GetInventoryCheck() *InventoryCheckDto {
	return r._inventoryCheck
}

var poolTaobaoInventoryMerchantAdjustAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoInventoryMerchantAdjustRequest()
	},
}

// GetTaobaoInventoryMerchantAdjustRequest 从 sync.Pool 获取 TaobaoInventoryMerchantAdjustAPIRequest
func GetTaobaoInventoryMerchantAdjustAPIRequest() *TaobaoInventoryMerchantAdjustAPIRequest {
	return poolTaobaoInventoryMerchantAdjustAPIRequest.Get().(*TaobaoInventoryMerchantAdjustAPIRequest)
}

// ReleaseTaobaoInventoryMerchantAdjustAPIRequest 将 TaobaoInventoryMerchantAdjustAPIRequest 放入 sync.Pool
func ReleaseTaobaoInventoryMerchantAdjustAPIRequest(v *TaobaoInventoryMerchantAdjustAPIRequest) {
	v.Reset()
	poolTaobaoInventoryMerchantAdjustAPIRequest.Put(v)
}
