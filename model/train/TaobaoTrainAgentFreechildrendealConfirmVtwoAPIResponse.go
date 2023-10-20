package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse 免费儿童处理 API返回值
// taobao.train.agent.freechildrendeal.confirm.vtwo
//
// 免费儿童列表查询
type TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponseModel).Reset()
}

// TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponseModel is 免费儿童处理 成功返回结果
type TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_freechildrendeal_confirm_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse)
	},
}

// GetTaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse
func GetTaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse() *TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse {
	return poolTaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse.Get().(*TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse)
}

// ReleaseTaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse 将 TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse(v *TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponse.Put(v)
}
