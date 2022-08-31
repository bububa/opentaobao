package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardStatusUpdateAPIResponse 服务商反馈服务的执行情况 API返回值
// tmall.servicecenter.workcard.status.update
//
// 1 如果服务商受理了此服务，修改合同状态为：已受理=3&lt;br/&gt;2 如果服务商没有受理此服务，修改合同状态为：已拒绝=10&lt;br/&gt;3 如果服务商执行了此服务，修改合同状态为：已执行=4&lt;br/&gt;4 如果服务商执行服务成功，修改合同状态为：已完成=5&lt;br/&gt;5 如果此工单为合同类型的工单，当服务商受理了此服务后，会进行分账
type TmallServicecenterWorkcardStatusUpdateAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardStatusUpdateAPIResponseModel
}

// TmallServicecenterWorkcardStatusUpdateAPIResponseModel is 服务商反馈服务的执行情况 成功返回结果
type TmallServicecenterWorkcardStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorMsgInfo string `json:"error_msg_info,omitempty" xml:"error_msg_info,omitempty"`
	// 错误码
	ErrorCodeValue string `json:"error_code_value,omitempty" xml:"error_code_value,omitempty"`
	// 是否执行成功
	Value bool `json:"value,omitempty" xml:"value,omitempty"`
	// 返回结果
	Rs bool `json:"rs,omitempty" xml:"rs,omitempty"`
}
