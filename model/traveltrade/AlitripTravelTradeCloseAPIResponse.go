package traveltrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelTradeCloseAPIResponse 飞猪度假-订单关闭接口（快速退款） API返回值
// alitrip.travel.trade.close
//
// 卖家关单（快速退款接口），不支持二次预约商品的订单
type AlitripTravelTradeCloseAPIResponse struct {
	model.CommonResponse
	AlitripTravelTradeCloseAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelTradeCloseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelTradeCloseAPIResponseModel).Reset()
}

// AlitripTravelTradeCloseAPIResponseModel is 飞猪度假-订单关闭接口（快速退款） 成功返回结果
type AlitripTravelTradeCloseAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_trade_close_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 交易关闭是否成功
	FirstResult bool `json:"first_result,omitempty" xml:"first_result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelTradeCloseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FirstResult = false
}

var poolAlitripTravelTradeCloseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelTradeCloseAPIResponse)
	},
}

// GetAlitripTravelTradeCloseAPIResponse 从 sync.Pool 获取 AlitripTravelTradeCloseAPIResponse
func GetAlitripTravelTradeCloseAPIResponse() *AlitripTravelTradeCloseAPIResponse {
	return poolAlitripTravelTradeCloseAPIResponse.Get().(*AlitripTravelTradeCloseAPIResponse)
}

// ReleaseAlitripTravelTradeCloseAPIResponse 将 AlitripTravelTradeCloseAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelTradeCloseAPIResponse(v *AlitripTravelTradeCloseAPIResponse) {
	v.Reset()
	poolAlitripTravelTradeCloseAPIResponse.Put(v)
}
