package legalsuit

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitCourttimePushAPIResponse 开庭时间推送（带附件） API返回值
// alibaba.legal.suit.courttime.push
//
// 开庭时间推送（带附件）
type AlibabaLegalSuitCourttimePushAPIResponse struct {
	model.CommonResponse
	AlibabaLegalSuitCourttimePushAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalSuitCourttimePushAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalSuitCourttimePushAPIResponseModel).Reset()
}

// AlibabaLegalSuitCourttimePushAPIResponseModel is 开庭时间推送（带附件） 成功返回结果
type AlibabaLegalSuitCourttimePushAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_suit_courttime_push_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误code
	ErrorCodeRes string `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 无返回值
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 错误消息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回是否成功标志
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalSuitCourttimePushAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrorCodeRes = ""
	m.Content = ""
	m.ErrorMsg = ""
	m.SuccessRes = false
}

var poolAlibabaLegalSuitCourttimePushAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalSuitCourttimePushAPIResponse)
	},
}

// GetAlibabaLegalSuitCourttimePushAPIResponse 从 sync.Pool 获取 AlibabaLegalSuitCourttimePushAPIResponse
func GetAlibabaLegalSuitCourttimePushAPIResponse() *AlibabaLegalSuitCourttimePushAPIResponse {
	return poolAlibabaLegalSuitCourttimePushAPIResponse.Get().(*AlibabaLegalSuitCourttimePushAPIResponse)
}

// ReleaseAlibabaLegalSuitCourttimePushAPIResponse 将 AlibabaLegalSuitCourttimePushAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalSuitCourttimePushAPIResponse(v *AlibabaLegalSuitCourttimePushAPIResponse) {
	v.Reset()
	poolAlibabaLegalSuitCourttimePushAPIResponse.Put(v)
}
