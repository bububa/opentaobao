package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】资源元素管理接口 APIResponse
taobao.alitrip.travel.item.element.manage

资源元素管理接口：提供商家管理（增删改）基本资源元素信息。基本资源元素可供多个商品共享
*/
type TaobaoAlitripTravelItemElementManageAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoAlitripTravelItemElementManageResponse `json:"alitrip_travel_item_element_manage_response,omitempty"` 
    TaobaoAlitripTravelItemElementManageResponse
}

/* model for simplify = false
type TaobaoAlitripTravelItemElementManageResponse struct {

    // firstResult
    
    FirstResult  *struct {
        TopElementResult  *TopElementResult `json:"top_element_result,omitempty"`
    } `json:"first_result,omitempty"`
    

}
*/

type TaobaoAlitripTravelItemElementManageResponse struct {

    // firstResult
    FirstResult   *TopElementResult `json:"first_result,omitempty"`

}
