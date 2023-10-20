package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicecenterworkcardserviceprogressupdateAPIResponse 更新服务进度 API返回值
// alibaba.servicecenter.workcard.serviceprogress.update
//
// 提供给外部合作服务商更新服务进度的接口
type AlibabaservicecenterworkcardserviceprogressupdateAPIResponse struct {
	model.CommonResponse
	AlibabaservicecenterworkcardserviceprogressupdateAPIResponseModel
}

// AlibabaservicecenterworkcardserviceprogressupdateAPIResponseModel is 更新服务进度 成功返回结果
type AlibabaservicecenterworkcardserviceprogressupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_servicecenter_workcard_serviceprogress_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
