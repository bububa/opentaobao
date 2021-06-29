package travel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
【API3.0】资源元素查询接口 APIResponse
taobao.alitrip.travel.item.element.query

提供资源元素查询接口，支持商家查询已经发布过的资源元素
*/
type TaobaoAlitripTravelItemElementQueryAPIResponse struct {
    model.CommonResponse
    TaobaoAlitripTravelItemElementQueryResponse
}

type TaobaoAlitripTravelItemElementQueryResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_item_element_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 资源元素列表
    
    Results   []TopElementParam `json:"results,omitempty" xml:"results>top_element_param,omitempty"`
    
    
}
