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
    Response *AliexpressSocialItemSearchResponse `json:"aliexpress_social_item_search_response,omitempty"`
}

type AliexpressSocialItemSearchResponse struct {

    // 报类型
    Result   *ItemPickPagingResult `json:"result,omitempty"`

}
