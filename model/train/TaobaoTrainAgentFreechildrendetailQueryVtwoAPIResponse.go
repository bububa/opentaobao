package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentfreechildrendetailqueryvtwoAPIResponse 免费儿童详情 API返回值
// taobao.train.agent.freechildrendetail.query.vtwo
//
// 免费儿童列表详情
type TaobaotrainagentfreechildrendetailqueryvtwoAPIResponse struct {
	model.CommonResponse
	TaobaotrainagentfreechildrendetailqueryvtwoAPIResponseModel
}

// TaobaotrainagentfreechildrendetailqueryvtwoAPIResponseModel is 免费儿童详情 成功返回结果
type TaobaotrainagentfreechildrendetailqueryvtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_freechildrendetail_query_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}
