package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫商品/SKU库存更新接口 APIResponse
tmall.item.quantity.update

天猫商品/SKU库存更新接口；支持商品库存更新；支持同一商品下的SKU批量更新。
*/
type TmallItemQuantityUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemQuantityUpdateResponse `json:"tmall_item_quantity_update_response,omitempty"` 
    TmallItemQuantityUpdateResponse
}

/* model for simplify = false
type TmallItemQuantityUpdateResponse struct {

    // 库存更新结果，商品id
    
    QuantityUpdateResult   string `json:"quantity_update_result,omitempty"`
    

}
*/

type TmallItemQuantityUpdateResponse struct {

    // 库存更新结果，商品id
    QuantityUpdateResult   string `json:"quantity_update_result,omitempty"`

}
