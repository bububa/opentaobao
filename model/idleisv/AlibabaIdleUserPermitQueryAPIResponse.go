package idleisv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleUserPermitQueryAPIResponse 查询服务商与卖家之间的订单消息绑定关系 API返回值
// alibaba.idle.user.permit.query
//
// 查询服务商与卖家之间的订单消息绑定关系
type AlibabaIdleUserPermitQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleUserPermitQueryAPIResponseModel
}

// AlibabaIdleUserPermitQueryAPIResponseModel is 查询服务商与卖家之间的订单消息绑定关系 成功返回结果
type AlibabaIdleUserPermitQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_user_permit_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultErrorCode string `json:"result_error_code,omitempty" xml:"result_error_code,omitempty"`
	// 错误信息
	ResultErrorMsg string `json:"result_error_msg,omitempty" xml:"result_error_msg,omitempty"`
	// 当前绑定状态，true为已经绑定，false为没有绑定
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 处理结果
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
