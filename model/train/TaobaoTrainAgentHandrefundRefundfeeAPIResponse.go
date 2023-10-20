package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentHandrefundRefundfeeAPIResponse 代理商手动退款接口 API返回值
// taobao.train.agent.handrefund.refundfee
//
// 火车票代理商手动退款接口
type TaobaoTrainAgentHandrefundRefundfeeAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentHandrefundRefundfeeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentHandrefundRefundfeeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentHandrefundRefundfeeAPIResponseModel).Reset()
}

// TaobaoTrainAgentHandrefundRefundfeeAPIResponseModel is 代理商手动退款接口 成功返回结果
type TaobaoTrainAgentHandrefundRefundfeeAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_handrefund_refundfee_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 失败code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 失败文案
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 是否成功标记
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentHandrefundRefundfeeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolTaobaoTrainAgentHandrefundRefundfeeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentHandrefundRefundfeeAPIResponse)
	},
}

// GetTaobaoTrainAgentHandrefundRefundfeeAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentHandrefundRefundfeeAPIResponse
func GetTaobaoTrainAgentHandrefundRefundfeeAPIResponse() *TaobaoTrainAgentHandrefundRefundfeeAPIResponse {
	return poolTaobaoTrainAgentHandrefundRefundfeeAPIResponse.Get().(*TaobaoTrainAgentHandrefundRefundfeeAPIResponse)
}

// ReleaseTaobaoTrainAgentHandrefundRefundfeeAPIResponse 将 TaobaoTrainAgentHandrefundRefundfeeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentHandrefundRefundfeeAPIResponse(v *TaobaoTrainAgentHandrefundRefundfeeAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentHandrefundRefundfeeAPIResponse.Put(v)
}
