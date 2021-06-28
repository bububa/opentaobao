package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建商家仓或者更新商家仓信息 APIResponse
taobao.inventory.warehouse.manage

创建商家仓或者更新商家仓信息
*/
type TaobaoInventoryWarehouseManageAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoInventoryWarehouseManageResponse `json:"inventory_warehouse_manage_response,omitempty"` 
    TaobaoInventoryWarehouseManageResponse
}

/* model for simplify = false
type TaobaoInventoryWarehouseManageResponse struct {

    // result
    
    Result  *struct {
        TaobaoInventoryWarehouseManageResult  *TaobaoInventoryWarehouseManageResult `json:"taobao_inventory_warehouse_manage_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoInventoryWarehouseManageResponse struct {

    // result
    Result   *TaobaoInventoryWarehouseManageResult `json:"result,omitempty"`

}
