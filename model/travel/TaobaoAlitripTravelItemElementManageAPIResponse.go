package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】资源元素管理接口 API返回值 
taobao.alitrip.travel.item.element.manage

资源元素管理接口：提供商家管理（增删改）基本资源元素信息。基本资源元素可供多个商品共享
*/
type TaobaoAlitripTravelItemElementManageAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelItemElementManageAPIResponseModel
}

// 【API3.0】资源元素管理接口 成功返回结果
type TaobaoAlitripTravelItemElementManageAPIResponseModel struct {
    XMLName xml.Name `xml:"alitrip_travel_item_element_manage_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // firstResult
    FirstResult   *TopElementResult `json:"first_result,omitempty" xml:"first_result,omitempty"`
}
