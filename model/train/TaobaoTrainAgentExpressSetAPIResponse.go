package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentExpressSetAPIResponse 线下票回填物流信息 API返回值
// taobao.train.agent.express.set
//
// 线下票回填物流信息服务
type TaobaoTrainAgentExpressSetAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentExpressSetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentExpressSetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentExpressSetAPIResponseModel).Reset()
}

// TaobaoTrainAgentExpressSetAPIResponseModel is 线下票回填物流信息 成功返回结果
type TaobaoTrainAgentExpressSetAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_express_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ErrorMsgCode string `json:"error_msg_code,omitempty" xml:"error_msg_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 扩展参数
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentExpressSetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsgCode = ""
	m.ErrorMsg = ""
	m.ExtendParams = ""
	m.IsSuccess = false
}

var poolTaobaoTrainAgentExpressSetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentExpressSetAPIResponse)
	},
}

// GetTaobaoTrainAgentExpressSetAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentExpressSetAPIResponse
func GetTaobaoTrainAgentExpressSetAPIResponse() *TaobaoTrainAgentExpressSetAPIResponse {
	return poolTaobaoTrainAgentExpressSetAPIResponse.Get().(*TaobaoTrainAgentExpressSetAPIResponse)
}

// ReleaseTaobaoTrainAgentExpressSetAPIResponse 将 TaobaoTrainAgentExpressSetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentExpressSetAPIResponse(v *TaobaoTrainAgentExpressSetAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentExpressSetAPIResponse.Put(v)
}
