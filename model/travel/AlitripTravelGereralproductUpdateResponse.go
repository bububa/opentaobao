package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
通用类目产品发布编辑 APIResponse
alitrip.travel.gereralproduct.update

提供给飞猪供销平台供应商发布编辑通用类目产品的API
*/
type AlitripTravelGereralproductUpdateAPIResponse struct {
    model.CommonResponse
    // Response *AlitripTravelGereralproductUpdateResponse `json:"alitrip_travel_gereralproduct_update_response,omitempty"` 
    AlitripTravelGereralproductUpdateResponse
}

/* model for simplify = false
type AlitripTravelGereralproductUpdateResponse struct {

    // firstResult
    
    FirstResult  *struct {
        TopTravelItem  *TopTravelItem `json:"top_travel_item,omitempty"`
    } `json:"first_result,omitempty"`
    

}
*/

type AlitripTravelGereralproductUpdateResponse struct {

    // firstResult
    FirstResult   *TopTravelItem `json:"first_result,omitempty"`

}
