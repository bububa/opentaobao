package train

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoTrainStopAgentCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainStopAgentCallbackAPIResponseModel).Reset()
}

// TaobaoTrainStopAgentCallbackAPIResponseModel is 火车票车次停运信息商家回调 成功返回结果
type TaobaoTrainStopAgentCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"train_stop_agent_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainStopAgentCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolTaobaoTrainStopAgentCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainStopAgentCallbackAPIResponse)
	},
}

// GetTaobaoTrainStopAgentCallbackAPIResponse 从 sync.Pool 获取 TaobaoTrainStopAgentCallbackAPIResponse
func GetTaobaoTrainStopAgentCallbackAPIResponse() *TaobaoTrainStopAgentCallbackAPIResponse {
	return poolTaobaoTrainStopAgentCallbackAPIResponse.Get().(*TaobaoTrainStopAgentCallbackAPIResponse)
}

// ReleaseTaobaoTrainStopAgentCallbackAPIResponse 将 TaobaoTrainStopAgentCallbackAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainStopAgentCallbackAPIResponse(v *TaobaoTrainStopAgentCallbackAPIResponse) {
	v.Reset()
	poolTaobaoTrainStopAgentCallbackAPIResponse.Put(v)
}
