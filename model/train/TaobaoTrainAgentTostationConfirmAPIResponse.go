package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentTostationConfirmAPIResponse 线下票确认送票至车站服务 API返回值
// taobao.train.agent.tostation.confirm
//
// 送票至车站的订单，代理商确认配送到站
type TaobaoTrainAgentTostationConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentTostationConfirmAPIResponseModel
}

// TaobaoTrainAgentTostationConfirmAPIResponseModel is 线下票确认送票至车站服务 成功返回结果
type TaobaoTrainAgentTostationConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_tostation_confirm_response"`
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
