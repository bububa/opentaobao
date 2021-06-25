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
    Response *TaobaoInventoryWarehouseManageResponse `json:"taobao_inventory_warehouse_manage_response,omitempty"`
}

type TaobaoInventoryWarehouseManageResponse struct {

    // result
    Result   *TaobaoInventoryWarehouseManageResult `json:"result,omitempty"`

}
