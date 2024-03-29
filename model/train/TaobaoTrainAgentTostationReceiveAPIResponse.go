package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentTostationReceiveAPIResponse 线下票送票至车站代理商确认用户已取票服务 API返回值
// taobao.train.agent.tostation.receive
//
// 送票至车站的订单，代理商确认用户已取票
type TaobaoTrainAgentTostationReceiveAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentTostationReceiveAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentTostationReceiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentTostationReceiveAPIResponseModel).Reset()
}

// TaobaoTrainAgentTostationReceiveAPIResponseModel is 线下票送票至车站代理商确认用户已取票服务 成功返回结果
type TaobaoTrainAgentTostationReceiveAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_tostation_receive_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ErrorMsgCode string `json:"error_msg_code,omitempty" xml:"error_msg_code,omitempty"`
	// 扩展参数
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentTostationReceiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsgCode = ""
	m.ExtendParams = ""
	m.ErrorMsg = ""
	m.IsSuccess = false
}

var poolTaobaoTrainAgentTostationReceiveAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentTostationReceiveAPIResponse)
	},
}

// GetTaobaoTrainAgentTostationReceiveAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentTostationReceiveAPIResponse
func GetTaobaoTrainAgentTostationReceiveAPIResponse() *TaobaoTrainAgentTostationReceiveAPIResponse {
	return poolTaobaoTrainAgentTostationReceiveAPIResponse.Get().(*TaobaoTrainAgentTostationReceiveAPIResponse)
}

// ReleaseTaobaoTrainAgentTostationReceiveAPIResponse 将 TaobaoTrainAgentTostationReceiveAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentTostationReceiveAPIResponse(v *TaobaoTrainAgentTostationReceiveAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentTostationReceiveAPIResponse.Put(v)
}
