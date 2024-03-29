package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentChangeissueConfirmVtwoAPIResponse 火车票代理商接口-跑腿改签出票回填-含鉴权校验 API返回值
// taobao.train.agent.changeissue.confirm.vtwo
//
// 火车票代理商接口-跑腿改签出票回填-含鉴权校验
type TaobaoTrainAgentChangeissueConfirmVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentChangeissueConfirmVtwoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentChangeissueConfirmVtwoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentChangeissueConfirmVtwoAPIResponseModel).Reset()
}

// TaobaoTrainAgentChangeissueConfirmVtwoAPIResponseModel is 火车票代理商接口-跑腿改签出票回填-含鉴权校验 成功返回结果
type TaobaoTrainAgentChangeissueConfirmVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_changeissue_confirm_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回填失败原因
	FailMessage string `json:"fail_message,omitempty" xml:"fail_message,omitempty"`
	// 回填失败code
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 出票回填结果
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentChangeissueConfirmVtwoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.FailMessage = ""
	m.FailCode = ""
	m.IsSuccess = false
}

var poolTaobaoTrainAgentChangeissueConfirmVtwoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentChangeissueConfirmVtwoAPIResponse)
	},
}

// GetTaobaoTrainAgentChangeissueConfirmVtwoAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentChangeissueConfirmVtwoAPIResponse
func GetTaobaoTrainAgentChangeissueConfirmVtwoAPIResponse() *TaobaoTrainAgentChangeissueConfirmVtwoAPIResponse {
	return poolTaobaoTrainAgentChangeissueConfirmVtwoAPIResponse.Get().(*TaobaoTrainAgentChangeissueConfirmVtwoAPIResponse)
}

// ReleaseTaobaoTrainAgentChangeissueConfirmVtwoAPIResponse 将 TaobaoTrainAgentChangeissueConfirmVtwoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentChangeissueConfirmVtwoAPIResponse(v *TaobaoTrainAgentChangeissueConfirmVtwoAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentChangeissueConfirmVtwoAPIResponse.Put(v)
}
