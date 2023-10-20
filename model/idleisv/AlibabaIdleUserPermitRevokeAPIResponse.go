package idleisv

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleUserPermitRevokeAPIResponse 删除服务商与卖家之间的订单消息绑定关系 API返回值
// alibaba.idle.user.permit.revoke
//
// 删除服务商与卖家之间的订单消息绑定关系，删除后不再发送消息
type AlibabaIdleUserPermitRevokeAPIResponse struct {
	model.CommonResponse
	AlibabaIdleUserPermitRevokeAPIResponseModel
}

// AlibabaIdleUserPermitRevokeAPIResponseModel is 删除服务商与卖家之间的订单消息绑定关系 成功返回结果
type AlibabaIdleUserPermitRevokeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_user_permit_revoke_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	ResultErrorCode string `json:"result_error_code,omitempty" xml:"result_error_code,omitempty"`
	// 错误信息
	ResultErrorMsg string `json:"result_error_msg,omitempty" xml:"result_error_msg,omitempty"`
	// 处理结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 处理结果
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
