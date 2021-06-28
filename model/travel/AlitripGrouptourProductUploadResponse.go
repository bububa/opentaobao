package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新版跟团游商品维护接口 APIResponse
alitrip.grouptour.product.upload

新版跟团游商品维护接口
*/
type AlitripGrouptourProductUploadAPIResponse struct {
    model.CommonResponse
    // Response *AlitripGrouptourProductUploadResponse `json:"alitrip_grouptour_product_upload_response,omitempty"` 
    AlitripGrouptourProductUploadResponse
}

/* model for simplify = false
type AlitripGrouptourProductUploadResponse struct {

    // firstResult
    
    FirstResult  *struct {
        TopTravelItem  *TopTravelItem `json:"top_travel_item,omitempty"`
    } `json:"first_result,omitempty"`
    

}
*/

type AlitripGrouptourProductUploadResponse struct {

    // firstResult
    FirstResult   *TopTravelItem `json:"first_result,omitempty"`

}
