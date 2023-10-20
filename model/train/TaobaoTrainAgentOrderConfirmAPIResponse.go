package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderConfirmAPIResponse 确认出票 API返回值
// taobao.train.agent.order.confirm
//
// 确认出票
type TaobaoTrainAgentOrderConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentOrderConfirmAPIResponseModel
}

// TaobaoTrainAgentOrderConfirmAPIResponseModel is 确认出票 成功返回结果
type TaobaoTrainAgentOrderConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_order_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}
