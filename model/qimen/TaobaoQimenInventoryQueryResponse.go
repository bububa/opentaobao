package qimen

import (
    "github.com/bububa/opentaobao/model"
)

/* 
库存查询接口（多商品） APIResponse
taobao.qimen.inventory.query

ERP调用奇门的接口,查询商品的库存量
*/
type TaobaoQimenInventoryQueryAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQimenInventoryQueryResponse `json:"qimen_inventory_query_response,omitempty"` 
    TaobaoQimenInventoryQueryResponse
}

/* model for simplify = false
type TaobaoQimenInventoryQueryResponse struct {

    // 
    
    Response  *struct {
        InventoryQueryResponse  *InventoryQueryResponse `json:"inventory_query_response,omitempty"`
    } `json:"response,omitempty"`
    

}
*/

type TaobaoQimenInventoryQueryResponse struct {

    // 
    Response   *InventoryQueryResponse `json:"response,omitempty"`

}
