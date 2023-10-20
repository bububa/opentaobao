package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderQueryAPIResponse 订单详情查询 API返回值
// taobao.train.agent.order.query
//
// 订单详情查询接口
type TaobaoTrainAgentOrderQueryAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentOrderQueryAPIResponseModel
}

// TaobaoTrainAgentOrderQueryAPIResponseModel is 订单详情查询 成功返回结果
type TaobaoTrainAgentOrderQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_order_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}
