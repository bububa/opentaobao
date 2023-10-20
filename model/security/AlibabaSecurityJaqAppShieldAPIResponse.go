package security

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabasecurityjaqappshieldAPIResponse 应用加固接口 API返回值
// alibaba.security.jaq.app.shield
//
// 提交应用进行应用加固,加固后需通过alibaba.security.jaq.app.shieldresult.get接口查询加固结果
type AlibabasecurityjaqappshieldAPIResponse struct {
	model.CommonResponse
	AlibabasecurityjaqappshieldAPIResponseModel
}

// AlibabasecurityjaqappshieldAPIResponseModel is 应用加固接口 成功返回结果
type AlibabasecurityjaqappshieldAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_security_jaq_app_shield_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 加固任务信息
	Result *ScanTaskInfo `json:"result,omitempty" xml:"result,omitempty"`
}
