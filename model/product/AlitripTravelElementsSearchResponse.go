package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家元素搜索 APIResponse
alitrip.travel.elements.search

提供商家维护的景点、酒店、餐饮等元素搜索
*/
type AlitripTravelElementsSearchAPIResponse struct {
    model.CommonResponse
    AlitripTravelElementsSearchResponse
}

type AlitripTravelElementsSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_elements_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象
    
    Result   *ResourceData `json:"result,omitempty" xml:"result,omitempty"`

    
}
