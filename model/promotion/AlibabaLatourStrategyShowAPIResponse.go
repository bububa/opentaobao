package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLatourStrategyShowAPIResponse 阿里巴巴权益投放接口 API返回值
// alibaba.latour.strategy.show
//
// 阿里巴巴权益平台权益投放接口
type AlibabaLatourStrategyShowAPIResponse struct {
	model.CommonResponse
	AlibabaLatourStrategyShowAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLatourStrategyShowAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLatourStrategyShowAPIResponseModel).Reset()
}

// AlibabaLatourStrategyShowAPIResponseModel is 阿里巴巴权益投放接口 成功返回结果
type AlibabaLatourStrategyShowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_latour_strategy_show_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaLatourStrategyShowResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLatourStrategyShowAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLatourStrategyShowAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLatourStrategyShowAPIResponse)
	},
}

// GetAlibabaLatourStrategyShowAPIResponse 从 sync.Pool 获取 AlibabaLatourStrategyShowAPIResponse
func GetAlibabaLatourStrategyShowAPIResponse() *AlibabaLatourStrategyShowAPIResponse {
	return poolAlibabaLatourStrategyShowAPIResponse.Get().(*AlibabaLatourStrategyShowAPIResponse)
}

// ReleaseAlibabaLatourStrategyShowAPIResponse 将 AlibabaLatourStrategyShowAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLatourStrategyShowAPIResponse(v *AlibabaLatourStrategyShowAPIResponse) {
	v.Reset()
	poolAlibabaLatourStrategyShowAPIResponse.Put(v)
}
