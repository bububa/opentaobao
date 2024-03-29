package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAccountStatusGetAPIResponse 查询账户级别关键词推广状态 API返回值
// alibaba.scbp.account.status.get
//
// 查询账户级别关键词推广状态
type AlibabaScbpAccountStatusGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAccountStatusGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAccountStatusGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAccountStatusGetAPIResponseModel).Reset()
}

// AlibabaScbpAccountStatusGetAPIResponseModel is 查询账户级别关键词推广状态 成功返回结果
type AlibabaScbpAccountStatusGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_account_status_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true:推广中,false:暂停
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAccountStatusGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaScbpAccountStatusGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAccountStatusGetAPIResponse)
	},
}

// GetAlibabaScbpAccountStatusGetAPIResponse 从 sync.Pool 获取 AlibabaScbpAccountStatusGetAPIResponse
func GetAlibabaScbpAccountStatusGetAPIResponse() *AlibabaScbpAccountStatusGetAPIResponse {
	return poolAlibabaScbpAccountStatusGetAPIResponse.Get().(*AlibabaScbpAccountStatusGetAPIResponse)
}

// ReleaseAlibabaScbpAccountStatusGetAPIResponse 将 AlibabaScbpAccountStatusGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAccountStatusGetAPIResponse(v *AlibabaScbpAccountStatusGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpAccountStatusGetAPIResponse.Put(v)
}
