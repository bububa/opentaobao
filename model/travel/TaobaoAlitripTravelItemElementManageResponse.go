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
    Response *TaobaoAlitripTravelItemElementManageResponse `json:"taobao_alitrip_travel_item_element_manage_response,omitempty"`
}

type TaobaoAlitripTravelItemElementManageResponse struct {

    // firstResult
    FirstResult   *TopElementResult `json:"first_result,omitempty"`

}
