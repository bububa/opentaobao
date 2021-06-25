package product

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家元素搜索 APIResponse
alitrip.travel.elements.search

提供商家维护的景点、酒店、餐饮等元素搜索
*/
type AlitripTravelElementsSearchAPIResponse struct {
    model.CommonResponse
    Response *AlitripTravelElementsSearchResponse `json:"alitrip_travel_elements_search_response,omitempty"`
}

type AlitripTravelElementsSearchResponse struct {

    // 返回对象
    Result   *ResourceData `json:"result,omitempty"`

}
