package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
天猫商品/SKU价格更新接口 APIResponse
tmall.item.price.update

天猫商品/SKU价格更新接口，支持商品、SKU价格同时更新，支持同一商品下的SKU批量更新。
*/
type TmallItemPriceUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TmallItemPriceUpdateResponse `json:"tmall_item_price_update_response,omitempty"` 
    TmallItemPriceUpdateResponse
}

/* model for simplify = false
type TmallItemPriceUpdateResponse struct {

    // 价格更新结果
    
    PriceUpdateResult   string `json:"price_update_result,omitempty"`
    

}
*/

type TmallItemPriceUpdateResponse struct {

    // 价格更新结果
    PriceUpdateResult   string `json:"price_update_result,omitempty"`

}
