package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentorderlockAPIResponse 锁单 API返回值
// taobao.train.agent.order.lock
//
// 锁单
type TaobaotrainagentorderlockAPIResponse struct {
	model.CommonResponse
	TaobaotrainagentorderlockAPIResponseModel
}

// TaobaotrainagentorderlockAPIResponseModel is 锁单 成功返回结果
type TaobaotrainagentorderlockAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_order_lock_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// rs
	Result *TapResult `json:"result,omitempty" xml:"result,omitempty"`
}
