package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse 更新服务进度 API返回值
// alibaba.servicecenter.workcard.serviceprogress.update
//
// 提供给外部合作服务商更新服务进度的接口
type AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponseModel).Reset()
}

// AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponseModel is 更新服务进度 成功返回结果
type AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponseModel struct {
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

// Reset 清空结构体
func (m *AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgCode = ""
	m.MsgInfo = ""
	m.IsSuccess = false
}

var poolAlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse)
	},
}

// GetAlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse 从 sync.Pool 获取 AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse
func GetAlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse() *AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse {
	return poolAlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse.Get().(*AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse)
}

// ReleaseAlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse 将 AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse(v *AlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse) {
	v.Reset()
	poolAlibabaServicecenterWorkcardServiceprogressUpdateAPIResponse.Put(v)
}
