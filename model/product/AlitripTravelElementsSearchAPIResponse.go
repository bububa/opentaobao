package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelElementsSearchAPIResponse 商家元素搜索 API返回值
// alitrip.travel.elements.search
//
// 提供商家维护的景点、酒店、餐饮等元素搜索
type AlitripTravelElementsSearchAPIResponse struct {
	model.CommonResponse
	AlitripTravelElementsSearchAPIResponseModel
}

// AlitripTravelElementsSearchAPIResponseModel is 商家元素搜索 成功返回结果
type AlitripTravelElementsSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_elements_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象
	Result *ResourceData `json:"result,omitempty" xml:"result,omitempty"`
}
