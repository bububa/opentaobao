package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
AE社交选品 APIResponse
aliexpress.social.item.search

AE社交选品,通过各种筛选条件对社交商品池进行筛选
*/
type AliexpressSocialItemSearchAPIResponse struct {
    model.CommonResponse
    // Response *AliexpressSocialItemSearchResponse `json:"aliexpress_social_item_search_response,omitempty"` 
    AliexpressSocialItemSearchResponse
}

/* model for simplify = false
type AliexpressSocialItemSearchResponse struct {

    // 报类型
    
    Result  *struct {
        ItemPickPagingResult  *ItemPickPagingResult `json:"item_pick_paging_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AliexpressSocialItemSearchResponse struct {

    // 报类型
    Result   *ItemPickPagingResult `json:"result,omitempty"`

}
