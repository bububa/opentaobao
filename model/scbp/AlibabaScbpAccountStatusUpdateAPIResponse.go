package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAccountStatusUpdateAPIResponse 修改账户级别关键词推广状态 API返回值
// alibaba.scbp.account.status.update
//
// 修改账户级别关键词推广状态
type AlibabaScbpAccountStatusUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAccountStatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAccountStatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAccountStatusUpdateAPIResponseModel).Reset()
}

// AlibabaScbpAccountStatusUpdateAPIResponseModel is 修改账户级别关键词推广状态 成功返回结果
type AlibabaScbpAccountStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_account_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAccountStatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaScbpAccountStatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAccountStatusUpdateAPIResponse)
	},
}

// GetAlibabaScbpAccountStatusUpdateAPIResponse 从 sync.Pool 获取 AlibabaScbpAccountStatusUpdateAPIResponse
func GetAlibabaScbpAccountStatusUpdateAPIResponse() *AlibabaScbpAccountStatusUpdateAPIResponse {
	return poolAlibabaScbpAccountStatusUpdateAPIResponse.Get().(*AlibabaScbpAccountStatusUpdateAPIResponse)
}

// ReleaseAlibabaScbpAccountStatusUpdateAPIResponse 将 AlibabaScbpAccountStatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAccountStatusUpdateAPIResponse(v *AlibabaScbpAccountStatusUpdateAPIResponse) {
	v.Reset()
	poolAlibabaScbpAccountStatusUpdateAPIResponse.Put(v)
}
