package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分页查询商家仓信息 APIResponse
taobao.inventory.warehouse.query

分页查询商家仓信息
*/
type TaobaoInventoryWarehouseQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoInventoryWarehouseQueryResponse `json:"inventory_warehouse_query_response,omitempty"` 
    TaobaoInventoryWarehouseQueryResponse
}

/* model for simplify = false
type TaobaoInventoryWarehouseQueryResponse struct {

    // result
    
    Result  *struct {
        PaginationResult  *PaginationResult `json:"pagination_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoInventoryWarehouseQueryResponse struct {

    // result
    Result   *PaginationResult `json:"result,omitempty"`

}
