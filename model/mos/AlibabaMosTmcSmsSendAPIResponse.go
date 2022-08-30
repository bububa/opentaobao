package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosTmcSmsSendAPIResponse 银泰TMC发送短信 API返回值
// alibaba.mos.tmc.sms.send
//
// 银泰TMC发送短信
type AlibabaMosTmcSmsSendAPIResponse struct {
	model.CommonResponse
	AlibabaMosTmcSmsSendAPIResponseModel
}

// AlibabaMosTmcSmsSendAPIResponseModel is 银泰TMC发送短信 成功返回结果
type AlibabaMosTmcSmsSendAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_tmc_sms_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	BizErrCode string `json:"biz_err_code,omitempty" xml:"biz_err_code,omitempty"`
	// 错误信息
	BizErrMessage string `json:"biz_err_message,omitempty" xml:"biz_err_message,omitempty"`
	// taskId
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
}
