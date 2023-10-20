package train

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentticketstatuscallbackAPIResponse 代理商票状态查询回调 API返回值
// taobao.train.agent.ticket.status.callback
//
// 代理商票状态查询结果回调
type TaobaotrainagentticketstatuscallbackAPIResponse struct {
	model.CommonResponse
	TaobaotrainagentticketstatuscallbackAPIResponseModel
}

// TaobaotrainagentticketstatuscallbackAPIResponseModel is 代理商票状态查询回调 成功返回结果
type TaobaotrainagentticketstatuscallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"train_agent_ticket_status_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功调用
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
