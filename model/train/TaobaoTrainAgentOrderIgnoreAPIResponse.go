package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderIgnoreAPIResponse 忽略订单 API返回值
// taobao.train.agent.order.ignore
//
// 忽略订单
type TaobaoTrainAgentOrderIgnoreAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentOrderIgnoreAPIResponseModel
}

// TaobaoTrainAgentOrderIgnoreAPIResponseModel is 忽略订单 成功返回结果
type TaobaoTrainAgentOrderIgnoreAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_order_ignore_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}
