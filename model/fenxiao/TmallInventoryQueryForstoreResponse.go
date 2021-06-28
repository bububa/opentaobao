package fenxiao

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询后端商品仓库库存 APIResponse
tmall.inventory.query.forstore

商家查询后端商品仓库库存
*/
type TmallInventoryQueryForstoreAPIResponse struct {
    model.CommonResponse
    // Response *TmallInventoryQueryForstoreResponse `json:"tmall_inventory_query_forstore_response,omitempty"` 
    TmallInventoryQueryForstoreResponse
}

/* model for simplify = false
type TmallInventoryQueryForstoreResponse struct {

    // 查询结果
    
    Result  *struct {
        InventoryQueryResult  *InventoryQueryResult `json:"inventory_query_result,omitempty"`
    } `json:"result,omitempty"`
    

    // 错误code
    
    Errorcode   string `json:"errorcode,omitempty"`
    

    // 错误信息
    
    Errormessage   string `json:"errormessage,omitempty"`
    

    // 整体成功或失败
    
    Issuccess   bool `json:"issuccess,omitempty"`
    

}
*/

type TmallInventoryQueryForstoreResponse struct {

    // 查询结果
    Result   *InventoryQueryResult `json:"result,omitempty"`

    // 错误code
    Errorcode   string `json:"errorcode,omitempty"`

    // 错误信息
    Errormessage   string `json:"errormessage,omitempty"`

    // 整体成功或失败
    Issuccess   bool `json:"issuccess,omitempty"`

}
