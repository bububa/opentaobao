package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentExpressSetVtwoAPIResponse 线下票回填物流信息v2--增加鉴权校验 API返回值
// taobao.train.agent.express.set.vtwo
//
// 线下票回填物流信息服务
type TaobaoTrainAgentExpressSetVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentExpressSetVtwoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentExpressSetVtwoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentExpressSetVtwoAPIResponseModel).Reset()
}

// TaobaoTrainAgentExpressSetVtwoAPIResponseModel is 线下票回填物流信息v2--增加鉴权校验 成功返回结果
type TaobaoTrainAgentExpressSetVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_express_set_vtwo_response"`
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
func (m *TaobaoTrainAgentExpressSetVtwoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorMsgCode = ""
	m.ErrorMsg = ""
	m.ExtendParams = ""
	m.IsSuccess = false
}

var poolTaobaoTrainAgentExpressSetVtwoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentExpressSetVtwoAPIResponse)
	},
}

// GetTaobaoTrainAgentExpressSetVtwoAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentExpressSetVtwoAPIResponse
func GetTaobaoTrainAgentExpressSetVtwoAPIResponse() *TaobaoTrainAgentExpressSetVtwoAPIResponse {
	return poolTaobaoTrainAgentExpressSetVtwoAPIResponse.Get().(*TaobaoTrainAgentExpressSetVtwoAPIResponse)
}

// ReleaseTaobaoTrainAgentExpressSetVtwoAPIResponse 将 TaobaoTrainAgentExpressSetVtwoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentExpressSetVtwoAPIResponse(v *TaobaoTrainAgentExpressSetVtwoAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentExpressSetVtwoAPIResponse.Put(v)
}
