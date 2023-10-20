package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkWholesaleOutboundorderCommitAPIResponse 盒马帮发货信息回传接口 API返回值
// alibaba.wdk.wholesale.outboundorder.commit
//
// 盒马帮发货信息回传接口
type AlibabaWdkWholesaleOutboundorderCommitAPIResponse struct {
	model.CommonResponse
	AlibabaWdkWholesaleOutboundorderCommitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkWholesaleOutboundorderCommitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkWholesaleOutboundorderCommitAPIResponseModel).Reset()
}

// AlibabaWdkWholesaleOutboundorderCommitAPIResponseModel is 盒马帮发货信息回传接口 成功返回结果
type AlibabaWdkWholesaleOutboundorderCommitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_wholesale_outboundorder_commit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkWholesaleOutboundorderCommitApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkWholesaleOutboundorderCommitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkWholesaleOutboundorderCommitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkWholesaleOutboundorderCommitAPIResponse)
	},
}

// GetAlibabaWdkWholesaleOutboundorderCommitAPIResponse 从 sync.Pool 获取 AlibabaWdkWholesaleOutboundorderCommitAPIResponse
func GetAlibabaWdkWholesaleOutboundorderCommitAPIResponse() *AlibabaWdkWholesaleOutboundorderCommitAPIResponse {
	return poolAlibabaWdkWholesaleOutboundorderCommitAPIResponse.Get().(*AlibabaWdkWholesaleOutboundorderCommitAPIResponse)
}

// ReleaseAlibabaWdkWholesaleOutboundorderCommitAPIResponse 将 AlibabaWdkWholesaleOutboundorderCommitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkWholesaleOutboundorderCommitAPIResponse(v *AlibabaWdkWholesaleOutboundorderCommitAPIResponse) {
	v.Reset()
	poolAlibabaWdkWholesaleOutboundorderCommitAPIResponse.Put(v)
}
