package category

import (
    "github.com/bububa/opentaobao/model"
)

/* 
展示类目获取接口 APIResponse
aliexpress.social.discategory.get

AE展示类目获取接口
*/
type AliexpressSocialDiscategoryGetAPIResponse struct {
    model.CommonResponse
    // Response *AliexpressSocialDiscategoryGetResponse `json:"aliexpress_social_discategory_get_response,omitempty"` 
    AliexpressSocialDiscategoryGetResponse
}

/* model for simplify = false
type AliexpressSocialDiscategoryGetResponse struct {

    // result
    
    Result  *struct {
        ItemPickPagingResult  *ItemPickPagingResult `json:"item_pick_paging_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AliexpressSocialDiscategoryGetResponse struct {

    // result
    Result   *ItemPickPagingResult `json:"result,omitempty"`

}
