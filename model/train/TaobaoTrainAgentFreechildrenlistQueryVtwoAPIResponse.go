package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse 免费儿童列表查询 API返回值
// taobao.train.agent.freechildrenlist.query.vtwo
//
// 免费儿童列表查询
type TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponseModel).Reset()
}

// TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponseModel is 免费儿童列表查询 成功返回结果
type TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_freechildrenlist_query_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse)
	},
}

// GetTaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse
func GetTaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse() *TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse {
	return poolTaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse.Get().(*TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse)
}

// ReleaseTaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse 将 TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse(v *TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponse.Put(v)
}
