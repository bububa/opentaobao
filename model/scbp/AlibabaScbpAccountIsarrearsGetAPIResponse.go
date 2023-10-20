package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAccountIsarrearsGetAPIResponse 查询关键词推广账户是否欠款 API返回值
// alibaba.scbp.account.isarrears.get
//
// 查询关键词推广账户是否欠款
type AlibabaScbpAccountIsarrearsGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAccountIsarrearsGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAccountIsarrearsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAccountIsarrearsGetAPIResponseModel).Reset()
}

// AlibabaScbpAccountIsarrearsGetAPIResponseModel is 查询关键词推广账户是否欠款 成功返回结果
type AlibabaScbpAccountIsarrearsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_account_isarrears_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 客户的关键词推广账户是否欠款
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAccountIsarrearsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaScbpAccountIsarrearsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAccountIsarrearsGetAPIResponse)
	},
}

// GetAlibabaScbpAccountIsarrearsGetAPIResponse 从 sync.Pool 获取 AlibabaScbpAccountIsarrearsGetAPIResponse
func GetAlibabaScbpAccountIsarrearsGetAPIResponse() *AlibabaScbpAccountIsarrearsGetAPIResponse {
	return poolAlibabaScbpAccountIsarrearsGetAPIResponse.Get().(*AlibabaScbpAccountIsarrearsGetAPIResponse)
}

// ReleaseAlibabaScbpAccountIsarrearsGetAPIResponse 将 AlibabaScbpAccountIsarrearsGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAccountIsarrearsGetAPIResponse(v *AlibabaScbpAccountIsarrearsGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpAccountIsarrearsGetAPIResponse.Put(v)
}
