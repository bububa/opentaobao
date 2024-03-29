package traveltrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTravelTradeServiceinfoWriteAPIResponse 订单服务信息写入接口 API返回值
// alitrip.travel.trade.serviceinfo.write
//
// 订单服务信息写入接口
type AlitripTravelTradeServiceinfoWriteAPIResponse struct {
	model.CommonResponse
	AlitripTravelTradeServiceinfoWriteAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTravelTradeServiceinfoWriteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTravelTradeServiceinfoWriteAPIResponseModel).Reset()
}

// AlitripTravelTradeServiceinfoWriteAPIResponseModel is 订单服务信息写入接口 成功返回结果
type AlitripTravelTradeServiceinfoWriteAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_trade_serviceinfo_write_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlitripTravelTradeServiceinfoWriteResultSet `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTravelTradeServiceinfoWriteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripTravelTradeServiceinfoWriteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTravelTradeServiceinfoWriteAPIResponse)
	},
}

// GetAlitripTravelTradeServiceinfoWriteAPIResponse 从 sync.Pool 获取 AlitripTravelTradeServiceinfoWriteAPIResponse
func GetAlitripTravelTradeServiceinfoWriteAPIResponse() *AlitripTravelTradeServiceinfoWriteAPIResponse {
	return poolAlitripTravelTradeServiceinfoWriteAPIResponse.Get().(*AlitripTravelTradeServiceinfoWriteAPIResponse)
}

// ReleaseAlitripTravelTradeServiceinfoWriteAPIResponse 将 AlitripTravelTradeServiceinfoWriteAPIResponse 保存到 sync.Pool
func ReleaseAlitripTravelTradeServiceinfoWriteAPIResponse(v *AlitripTravelTradeServiceinfoWriteAPIResponse) {
	v.Reset()
	poolAlitripTravelTradeServiceinfoWriteAPIResponse.Put(v)
}
