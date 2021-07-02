package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmCreateAPIResponse 创建客资 API返回值
// alibaba.lsy.crm.create
//
// 欢客调用该接口进行客资创建
type AlibabaLsyCrmCreateAPIResponse struct {
	model.CommonResponse
	AlibabaLsyCrmCreateAPIResponseModel
}

// AlibabaLsyCrmCreateAPIResponseModel is 创建客资 成功返回结果
type AlibabaLsyCrmCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lsy_crm_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误提示
	FailMsg string `json:"fail_msg,omitempty" xml:"fail_msg,omitempty"`
	// 错误码
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
	// 返回的数据
	Data *NrtCreateRecordReturnDto `json:"data,omitempty" xml:"data,omitempty"`
}
