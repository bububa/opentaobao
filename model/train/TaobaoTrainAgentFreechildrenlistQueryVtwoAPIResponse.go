package train

import (
	"encoding/xml"

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

// TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponseModel is 免费儿童列表查询 成功返回结果
type TaobaoTrainAgentFreechildrenlistQueryVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_freechildrenlist_query_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}
