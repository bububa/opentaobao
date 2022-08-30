package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentOrderLockAPIResponse 锁单 API返回值
// taobao.train.agent.order.lock
//
// 锁单
type TaobaoTrainAgentOrderLockAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentOrderLockAPIResponseModel
}

// TaobaoTrainAgentOrderLockAPIResponseModel is 锁单 成功返回结果
type TaobaoTrainAgentOrderLockAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_order_lock_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}
