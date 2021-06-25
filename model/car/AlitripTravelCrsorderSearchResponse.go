package car

import (
    "github.com/bububa/opentaobao/model"
)

/* 
CRS接送机订单列表搜索 APIResponse
alitrip.travel.crsorder.search

提供给CRS商家搜索订单列表，仅返回订单号列表
*/
type AlitripTravelCrsorderSearchAPIResponse struct {
    model.CommonResponse
    Response *AlitripTravelCrsorderSearchResponse `json:"alitrip_travel_crsorder_search_response,omitempty"`
}

type AlitripTravelCrsorderSearchResponse struct {

    // 订单id列表（string类型）
    OrderStringList   []String `json:"order_string_list,omitempty"`

}
