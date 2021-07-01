package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLsyCrmUpdateAPIResponse
跟进客资状态接口 API返回值
alibaba.lsy.crm.update

同步客资状态接口 */
type AlibabaLsyCrmUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaLsyCrmUpdateAPIResponseModel
}

// AlibabaLsyCrmUpdateAPIResponseModel is 跟进客资状态接口 成功返回结果
type AlibabaLsyCrmUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误提示
	FailMsg string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`
	// 错误码
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
	// 是否成功跟进客资状态
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}
