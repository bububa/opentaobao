package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse 免费儿童详情 API返回值
// taobao.train.agent.freechildrendetail.query.vtwo
//
// 免费儿童列表详情
type TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponseModel).Reset()
}

// TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponseModel is 免费儿童详情 成功返回结果
type TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_freechildrendetail_query_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse)
	},
}

// GetTaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse
func GetTaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse() *TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse {
	return poolTaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse.Get().(*TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse)
}

// ReleaseTaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse 将 TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse(v *TaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentFreechildrendetailQueryVtwoAPIResponse.Put(v)
}
