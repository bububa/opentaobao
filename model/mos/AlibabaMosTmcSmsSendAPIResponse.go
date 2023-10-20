package mos

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *AlibabaMosTmcSmsSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosTmcSmsSendAPIResponseModel).Reset()
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

// Reset 清空结构体
func (m *AlibabaMosTmcSmsSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BizErrCode = ""
	m.BizErrMessage = ""
	m.Data = 0
}

var poolAlibabaMosTmcSmsSendAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosTmcSmsSendAPIResponse)
	},
}

// GetAlibabaMosTmcSmsSendAPIResponse 从 sync.Pool 获取 AlibabaMosTmcSmsSendAPIResponse
func GetAlibabaMosTmcSmsSendAPIResponse() *AlibabaMosTmcSmsSendAPIResponse {
	return poolAlibabaMosTmcSmsSendAPIResponse.Get().(*AlibabaMosTmcSmsSendAPIResponse)
}

// ReleaseAlibabaMosTmcSmsSendAPIResponse 将 AlibabaMosTmcSmsSendAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosTmcSmsSendAPIResponse(v *AlibabaMosTmcSmsSendAPIResponse) {
	v.Reset()
	poolAlibabaMosTmcSmsSendAPIResponse.Put(v)
}
