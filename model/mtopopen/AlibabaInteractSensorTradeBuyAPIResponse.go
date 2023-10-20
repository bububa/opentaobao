package mtopopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorTradeBuyAPIResponse 手淘下单能力开放 API返回值
// alibaba.interact.sensor.trade.buy
//
// 交易流程鉴权
type AlibabaInteractSensorTradeBuyAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorTradeBuyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorTradeBuyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorTradeBuyAPIResponseModel).Reset()
}

// AlibabaInteractSensorTradeBuyAPIResponseModel is 手淘下单能力开放 成功返回结果
type AlibabaInteractSensorTradeBuyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_trade_buy_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorTradeBuyAPIResponseModel) Reset() {
	m.RequestId = ""
}

var poolAlibabaInteractSensorTradeBuyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorTradeBuyAPIResponse)
	},
}

// GetAlibabaInteractSensorTradeBuyAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorTradeBuyAPIResponse
func GetAlibabaInteractSensorTradeBuyAPIResponse() *AlibabaInteractSensorTradeBuyAPIResponse {
	return poolAlibabaInteractSensorTradeBuyAPIResponse.Get().(*AlibabaInteractSensorTradeBuyAPIResponse)
}

// ReleaseAlibabaInteractSensorTradeBuyAPIResponse 将 AlibabaInteractSensorTradeBuyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorTradeBuyAPIResponse(v *AlibabaInteractSensorTradeBuyAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorTradeBuyAPIResponse.Put(v)
}
