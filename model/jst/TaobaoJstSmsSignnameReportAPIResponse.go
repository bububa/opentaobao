package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsSignnameReportAPIResponse 服务商存量短信签名上传 API返回值
// taobao.jst.sms.signname.report
//
// 用于上传目前已经在阿里通信申请到的且正在使用的签名数据，确保签名数据正确，否则会导致短信发送失败！！！
type TaobaoJstSmsSignnameReportAPIResponse struct {
	model.CommonResponse
	TaobaoJstSmsSignnameReportAPIResponseModel
}

// TaobaoJstSmsSignnameReportAPIResponseModel is 服务商存量短信签名上传 成功返回结果
type TaobaoJstSmsSignnameReportAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_signname_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	RCode string `json:"r_code,omitempty" xml:"r_code,omitempty"`
	// 请求成功
	RSuccess string `json:"r_success,omitempty" xml:"r_success,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}
