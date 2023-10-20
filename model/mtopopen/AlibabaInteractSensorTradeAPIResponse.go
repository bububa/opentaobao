package mtopopen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractSensorTradeAPIResponse 交易组件 API返回值
// alibaba.interact.sensor.trade
//
// 交易流程
type AlibabaInteractSensorTradeAPIResponse struct {
	model.CommonResponse
	AlibabaInteractSensorTradeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaInteractSensorTradeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaInteractSensorTradeAPIResponseModel).Reset()
}

// AlibabaInteractSensorTradeAPIResponseModel is 交易组件 成功返回结果
type AlibabaInteractSensorTradeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_sensor_trade_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result=1
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaInteractSensorTradeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaInteractSensorTradeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaInteractSensorTradeAPIResponse)
	},
}

// GetAlibabaInteractSensorTradeAPIResponse 从 sync.Pool 获取 AlibabaInteractSensorTradeAPIResponse
func GetAlibabaInteractSensorTradeAPIResponse() *AlibabaInteractSensorTradeAPIResponse {
	return poolAlibabaInteractSensorTradeAPIResponse.Get().(*AlibabaInteractSensorTradeAPIResponse)
}

// ReleaseAlibabaInteractSensorTradeAPIResponse 将 AlibabaInteractSensorTradeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaInteractSensorTradeAPIResponse(v *AlibabaInteractSensorTradeAPIResponse) {
	v.Reset()
	poolAlibabaInteractSensorTradeAPIResponse.Put(v)
}
