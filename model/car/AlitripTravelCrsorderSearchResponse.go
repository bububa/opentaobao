package car

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
CRS接送机订单列表搜索 APIResponse
alitrip.travel.crsorder.search

提供给CRS商家搜索订单列表，仅返回订单号列表
*/
type AlitripTravelCrsorderSearchAPIResponse struct {
    model.CommonResponse
    AlitripTravelCrsorderSearchResponse
}

type AlitripTravelCrsorderSearchResponse struct {
    XMLName xml.Name `xml:"alitrip_travel_crsorder_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 订单id列表（string类型）
    
    OrderStringList   []string `json:"order_string_list,omitempty" xml:"order_string_list>string,omitempty"`
    
    
}
