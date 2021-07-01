package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentHoldseatConfirmAPIResponse
火车票代理商接口——确认占座是否成功 API返回值
taobao.train.agent.holdseat.confirm

火车票代理商接口——确认占座是否成功 */
type TaobaoTrainAgentHoldseatConfirmAPIResponse struct {
	model.CommonResponse
	TaobaoTrainAgentHoldseatConfirmAPIResponseModel
}

// TaobaoTrainAgentHoldseatConfirmAPIResponseModel is 火车票代理商接口——确认占座是否成功 成功返回结果
type TaobaoTrainAgentHoldseatConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_holdseat_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// resultMsg
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// resultCode
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
