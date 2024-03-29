package traveltrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelTradesSearchAPIResponse 飞猪度假-订单列表搜索接口 API返回值
// alitrip.travel.trades.search
//
// 订单列表搜索接口：以订单创建、结束时间、订单状态为搜索条件，搜索过滤出满足条件的卖家订单列表。
type AlitripTravelTradesSearchAPIResponse struct {
	model.CommonResponse
	AlitripTravelTradesSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelTradesSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelTradesSearchAPIResponseModel).Reset()
}

// AlitripTravelTradesSearchAPIResponseModel is 飞猪度假-订单列表搜索接口 成功返回结果
type AlitripTravelTradesSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_trades_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 主订单id列表（Long类型）
	OrderList []int64 `json:"order_list,omitempty" xml:"order_list>int64,omitempty"`
	// 主订单id列表（string类型）
	OrderStringList []string `json:"order_string_list,omitempty" xml:"order_string_list>string,omitempty"`
	// 本次搜索包含的订单总数，用于分页控制
	TotalOrders int64 `json:"total_orders,omitempty" xml:"total_orders,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelTradesSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderList = m.OrderList[:0]
	m.OrderStringList = m.OrderStringList[:0]
	m.TotalOrders = 0
}

var poolAlitripTravelTradesSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelTradesSearchAPIResponse)
	},
}

// GetAlitripTravelTradesSearchAPIResponse 从 sync.Pool 获取 AlitripTravelTradesSearchAPIResponse
func GetAlitripTravelTradesSearchAPIResponse() *AlitripTravelTradesSearchAPIResponse {
	return poolAlitripTravelTradesSearchAPIResponse.Get().(*AlitripTravelTradesSearchAPIResponse)
}

// ReleaseAlitripTravelTradesSearchAPIResponse 将 AlitripTravelTradesSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelTradesSearchAPIResponse(v *AlitripTravelTradesSearchAPIResponse) {
	v.Reset()
	poolAlitripTravelTradesSearchAPIResponse.Put(v)
}
