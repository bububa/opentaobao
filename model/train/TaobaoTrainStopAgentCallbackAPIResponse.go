package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainStopAgentCallbackAPIResponse 火车票车次停运信息商家回调 API返回值
// taobao.train.stop.agent.callback
//
// 火车票车次停运信息商家回调接口
type TaobaoTrainStopAgentCallbackAPIResponse struct {
	model.CommonResponse
	TaobaoTrainStopAgentCallbackAPIResponseModel
}

// TaobaoTrainStopAgentCallbackAPIResponseModel is 火车票车次停运信息商家回调 成功返回结果
type TaobaoTrainStopAgentCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"train_stop_agent_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
