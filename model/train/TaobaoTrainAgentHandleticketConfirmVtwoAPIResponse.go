package train

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentHandleticketConfirmVtwoAPIResponse 代理商出票中v2--增加鉴权校验 API返回值
// taobao.train.agent.handleticket.confirm.vtwo
//
// 代理商出票中
type TaobaoTrainAgentHandleticketConfirmVtwoAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentHandleticketConfirmVtwoAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTrainAgentHandleticketConfirmVtwoAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTrainAgentHandleticketConfirmVtwoAPIResponseModel).Reset()
}

// TaobaoTrainAgentHandleticketConfirmVtwoAPIResponseModel is 代理商出票中v2--增加鉴权校验 成功返回结果
type TaobaoTrainAgentHandleticketConfirmVtwoAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_handleticket_confirm_vtwo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	TrainErrorCode string `json:"train_error_code,omitempty" xml:"train_error_code,omitempty"`
	// 错误信息
	TrainErrorMsg string `json:"train_error_msg,omitempty" xml:"train_error_msg,omitempty"`
	// 暂无
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTrainAgentHandleticketConfirmVtwoAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TrainErrorCode = ""
	m.TrainErrorMsg = ""
	m.ExtendParams = ""
	m.IsSuccess = false
}

var poolTaobaoTrainAgentHandleticketConfirmVtwoAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTrainAgentHandleticketConfirmVtwoAPIResponse)
	},
}

// GetTaobaoTrainAgentHandleticketConfirmVtwoAPIResponse 从 sync.Pool 获取 TaobaoTrainAgentHandleticketConfirmVtwoAPIResponse
func GetTaobaoTrainAgentHandleticketConfirmVtwoAPIResponse() *TaobaoTrainAgentHandleticketConfirmVtwoAPIResponse {
	return poolTaobaoTrainAgentHandleticketConfirmVtwoAPIResponse.Get().(*TaobaoTrainAgentHandleticketConfirmVtwoAPIResponse)
}

// ReleaseTaobaoTrainAgentHandleticketConfirmVtwoAPIResponse 将 TaobaoTrainAgentHandleticketConfirmVtwoAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTrainAgentHandleticketConfirmVtwoAPIResponse(v *TaobaoTrainAgentHandleticketConfirmVtwoAPIResponse) {
	v.Reset()
	poolTaobaoTrainAgentHandleticketConfirmVtwoAPIResponse.Put(v)
}
