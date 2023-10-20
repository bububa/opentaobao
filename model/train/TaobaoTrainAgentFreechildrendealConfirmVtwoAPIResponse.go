package train

import (
	"encoding/xml"

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

// TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponseModel is 免费儿童处理 成功返回结果
type TaobaoTrainAgentFreechildrendealConfirmVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_freechildrendeal_confirm_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}
