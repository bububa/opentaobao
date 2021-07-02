package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderPayAPIResponse 代购订单代付接口 API返回值
// taobao.train.agent.order.pay
//
// 代购订单代付接口
type TaobaoTrainAgentOrderPayAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentOrderPayAPIResponseModel
}

// TaobaoTrainAgentOrderPayAPIResponseModel is 代购订单代付接口 成功返回结果
type TaobaoTrainAgentOrderPayAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_order_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功返回
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
}
