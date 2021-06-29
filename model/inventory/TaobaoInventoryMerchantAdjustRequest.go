package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
货品库存商家端调整 API请求
taobao.inventory.merchant.adjust

货品库存商家端调整 ，入库，出库，盘点
*/
type TaobaoInventoryMerchantAdjustRequest struct {
    model.Params
    // 调整库存对象
    _inventoryCheck   *InventoryCheckDTO
}

// 初始化TaobaoInventoryMerchantAdjustRequest对象
func NewTaobaoInventoryMerchantAdjustRequest() *TaobaoInventoryMerchantAdjustRequest{
    return &TaobaoInventoryMerchantAdjustRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoInventoryMerchantAdjustRequest) GetApiMethodName() string {
    return "taobao.inventory.merchant.adjust"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoInventoryMerchantAdjustRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// InventoryCheck Setter
// 调整库存对象
func (r *TaobaoInventoryMerchantAdjustRequest) SetInventoryCheck(_inventoryCheck *InventoryCheckDTO) error {
    r._inventoryCheck = _inventoryCheck
    r.Set("inventory_check", _inventoryCheck)
    return nil
}

// InventoryCheck Getter
func (r TaobaoInventoryMerchantAdjustRequest) GetInventoryCheck() *InventoryCheckDTO {
    return r._inventoryCheck
}
