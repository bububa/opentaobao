package legalcase

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseCommonNoticeAPIResponse 消息通知 API返回值
// alibaba.legal.case.common.notice
//
// 同步通知给诉讼系统
type AlibabaLegalCaseCommonNoticeAPIResponse struct {
	model.CommonResponse
	AlibabaLegalCaseCommonNoticeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLegalCaseCommonNoticeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLegalCaseCommonNoticeAPIResponseModel).Reset()
}

// AlibabaLegalCaseCommonNoticeAPIResponseModel is 消息通知 成功返回结果
type AlibabaLegalCaseCommonNoticeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_case_common_notice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// error
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// content
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// msg
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLegalCaseCommonNoticeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errcode = ""
	m.Content = ""
	m.Errmsg = ""
	m.IsSuccess = false
}

var poolAlibabaLegalCaseCommonNoticeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLegalCaseCommonNoticeAPIResponse)
	},
}

// GetAlibabaLegalCaseCommonNoticeAPIResponse 从 sync.Pool 获取 AlibabaLegalCaseCommonNoticeAPIResponse
func GetAlibabaLegalCaseCommonNoticeAPIResponse() *AlibabaLegalCaseCommonNoticeAPIResponse {
	return poolAlibabaLegalCaseCommonNoticeAPIResponse.Get().(*AlibabaLegalCaseCommonNoticeAPIResponse)
}

// ReleaseAlibabaLegalCaseCommonNoticeAPIResponse 将 AlibabaLegalCaseCommonNoticeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLegalCaseCommonNoticeAPIResponse(v *AlibabaLegalCaseCommonNoticeAPIResponse) {
	v.Reset()
	poolAlibabaLegalCaseCommonNoticeAPIResponse.Put(v)
}
