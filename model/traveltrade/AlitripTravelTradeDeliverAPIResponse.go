package traveltrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelTradeDeliverAPIResponse 飞猪度假-订单发货接口 API返回值
// alitrip.travel.trade.deliver
//
// 航旅度假无需物流普通商品发货接口（不支持二次预约商品），只支持子订单级别发货
type AlitripTravelTradeDeliverAPIResponse struct {
	model.CommonResponse
	AlitripTravelTradeDeliverAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelTradeDeliverAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelTradeDeliverAPIResponseModel).Reset()
}

// AlitripTravelTradeDeliverAPIResponseModel is 飞猪度假-订单发货接口 成功返回结果
type AlitripTravelTradeDeliverAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_trade_deliver_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 发货是否成功
	FirstResult bool `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelTradeDeliverAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FirstResult = false
}

var poolAlitripTravelTradeDeliverAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelTradeDeliverAPIResponse)
	},
}

// GetAlitripTravelTradeDeliverAPIResponse 从 sync.Pool 获取 AlitripTravelTradeDeliverAPIResponse
func GetAlitripTravelTradeDeliverAPIResponse() *AlitripTravelTradeDeliverAPIResponse {
	return poolAlitripTravelTradeDeliverAPIResponse.Get().(*AlitripTravelTradeDeliverAPIResponse)
}

// ReleaseAlitripTravelTradeDeliverAPIResponse 将 AlitripTravelTradeDeliverAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelTradeDeliverAPIResponse(v *AlitripTravelTradeDeliverAPIResponse) {
	v.Reset()
	poolAlitripTravelTradeDeliverAPIResponse.Put(v)
}
