package customizemarket

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallindustrybabyauthprofilebackflowAPIResponse 孕校云回流档案 API返回值
// tmall.industry.baby.authprofile.backflow
//
// 孕校云回流档案
type TmallindustrybabyauthprofilebackflowAPIResponse struct {
	model.CommonResponse
	TmallindustrybabyauthprofilebackflowAPIResponseModel
}

// TmallindustrybabyauthprofilebackflowAPIResponseModel is 孕校云回流档案 成功返回结果
type TmallindustrybabyauthprofilebackflowAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_industry_baby_authprofile_backflow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	MsgMessage string `json:"msg_message,omitempty" xml:"msg_message,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 接口调用是否成功
	ResultStatus bool `json:"result_status,omitempty" xml:"result_status,omitempty"`
}
