package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicecenterworkcardserviceskusuggestAPIResponse 服务商反馈需要履行的服务项 API返回值
// alibaba.servicecenter.workcard.servicesku.suggest
//
// 服务商反馈需要履行的服务项
type AlibabaservicecenterworkcardserviceskusuggestAPIResponse struct {
	model.CommonResponse
	AlibabaservicecenterworkcardserviceskusuggestAPIResponseModel
}

// AlibabaservicecenterworkcardserviceskusuggestAPIResponseModel is 服务商反馈需要履行的服务项 成功返回结果
type AlibabaservicecenterworkcardserviceskusuggestAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_workcard_servicesku_suggest_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
