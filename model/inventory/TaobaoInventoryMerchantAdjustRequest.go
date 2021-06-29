package inventory

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
货品库存商家端调整 APIRequest
taobao.inventory.merchant.adjust

货品库存商家端调整 ，入库，出库，盘点
*/
type TaobaoInventoryMerchantAdjustRequest struct {
    model.Params

    // 调整库存对象
    inventoryCheck   *InventoryCheckDto 

}

func NewTaobaoInventoryMerchantAdjustRequest() *TaobaoInventoryMerchantAdjustRequest{
    return &TaobaoInventoryMerchantAdjustRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoInventoryMerchantAdjustRequest) GetApiMethodName() string {
    return "taobao.inventory.merchant.adjust"
}

func (r TaobaoInventoryMerchantAdjustRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoInventoryMerchantAdjustRequest) SetInventoryCheck(inventoryCheck *InventoryCheckDto) error {
    r.inventoryCheck = inventoryCheck
    r.Set("inventory_check", inventoryCheck)
    return nil
}

func (r TaobaoInventoryMerchantAdjustRequest) GetInventoryCheck() *InventoryCheckDto {
    return r.inventoryCheck
}

