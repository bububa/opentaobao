package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentHandleticketConfirmAPIResponse 代理商出票中 API返回值
// taobao.train.agent.handleticket.confirm
//
// 代理商出票中
type TaobaoTrainAgentHandleticketConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentHandleticketConfirmAPIResponseModel
}

// TaobaoTrainAgentHandleticketConfirmAPIResponseModel is 代理商出票中 成功返回结果
type TaobaoTrainAgentHandleticketConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_handleticket_confirm_response"`
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
