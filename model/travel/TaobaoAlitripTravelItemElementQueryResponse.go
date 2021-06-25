package travel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】资源元素查询接口 APIResponse
taobao.alitrip.travel.item.element.query

提供资源元素查询接口，支持商家查询已经发布过的资源元素
*/
type TaobaoAlitripTravelItemElementQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAlitripTravelItemElementQueryResponse `json:"taobao_alitrip_travel_item_element_query_response,omitempty"`
}

type TaobaoAlitripTravelItemElementQueryResponse struct {

    // 资源元素列表
    Results   []TopElementParam `json:"results,omitempty"`

}
