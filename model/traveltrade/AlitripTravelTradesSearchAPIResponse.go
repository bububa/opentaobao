package traveltrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTravelTradesSearchAPIResponse
飞猪度假-订单列表搜索接口 API返回值
alitrip.travel.trades.search

订单列表搜索接口：以订单创建、结束时间、订单状态为搜索条件，搜索过滤出满足条件的卖家订单列表。 */
type AlitripTravelTradesSearchAPIResponse struct {
	model.CommonResponse
	AlitripTravelTradesSearchAPIResponseModel
}

// AlitripTravelTradesSearchAPIResponseModel is 飞猪度假-订单列表搜索接口 成功返回结果
type AlitripTravelTradesSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_trades_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 本次搜索包含的订单总数，用于分页控制
	TotalOrders int64 `json:"total_orders,omitempty" xml:"total_orders,omitempty"`
	// 主订单id列表（Long类型）
	OrderList []int64 `json:"order_list,omitempty" xml:"order_list>int64,omitempty"`
	// 主订单id列表（string类型）
	OrderStringList []string `json:"order_string_list,omitempty" xml:"order_string_list>string,omitempty"`
}
