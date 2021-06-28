package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】资源元素管理接口 APIResponse
taobao.alitrip.travel.item.element.manage

资源元素管理接口：提供商家管理（增删改）基本资源元素信息。基本资源元素可供多个商品共享
*/
type TaobaoAlitripTravelItemElementManageAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alitrip_travel_item_element_manage_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // firstResult
    
    FirstResult   *TopElementResult `json:"first_result,omitempty" xml:"