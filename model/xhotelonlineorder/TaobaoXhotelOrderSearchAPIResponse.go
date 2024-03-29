package xhotelonlineorder

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderSearchAPIResponse 酒店产品库订单查询 API返回值
// taobao.xhotel.order.search
//
// 酒店产品库订单查询功能，查询90天内的订单
type TaobaoXhotelOrderSearchAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelOrderSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelOrderSearchAPIResponseModel).Reset()
}

// TaobaoXhotelOrderSearchAPIResponseModel is 酒店产品库订单查询 成功返回结果
type TaobaoXhotelOrderSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_order_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单信息
	HotelOrders []XHotelOrder `json:"hotel_orders,omitempty" xml:"hotel_orders>x_hotel_order,omitempty"`
	// 符合条件的结果总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelOrderSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.HotelOrders = m.HotelOrders[:0]
	m.TotalResults = 0
}

var poolTaobaoXhotelOrderSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelOrderSearchAPIResponse)
	},
}

// GetTaobaoXhotelOrderSearchAPIResponse 从 sync.Pool 获取 TaobaoXhotelOrderSearchAPIResponse
func GetTaobaoXhotelOrderSearchAPIResponse() *TaobaoXhotelOrderSearchAPIResponse {
	return poolTaobaoXhotelOrderSearchAPIResponse.Get().(*TaobaoXhotelOrderSearchAPIResponse)
}

// ReleaseTaobaoXhotelOrderSearchAPIResponse 将 TaobaoXhotelOrderSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelOrderSearchAPIResponse(v *TaobaoXhotelOrderSearchAPIResponse) {
	v.Reset()
	poolTaobaoXhotelOrderSearchAPIResponse.Put(v)
}
