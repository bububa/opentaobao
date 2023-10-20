package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkWholesaleInboundorderCommitAPIResponse 盒马帮退货信息回传接口 API返回值
// alibaba.wdk.wholesale.inboundorder.commit
//
// 盒马帮退货信息回传接口
type AlibabaWdkWholesaleInboundorderCommitAPIResponse struct {
	model.CommonResponse
	AlibabaWdkWholesaleInboundorderCommitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkWholesaleInboundorderCommitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkWholesaleInboundorderCommitAPIResponseModel).Reset()
}

// AlibabaWdkWholesaleInboundorderCommitAPIResponseModel is 盒马帮退货信息回传接口 成功返回结果
type AlibabaWdkWholesaleInboundorderCommitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_wholesale_inboundorder_commit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkWholesaleInboundorderCommitApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkWholesaleInboundorderCommitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkWholesaleInboundorderCommitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkWholesaleInboundorderCommitAPIResponse)
	},
}

// GetAlibabaWdkWholesaleInboundorderCommitAPIResponse 从 sync.Pool 获取 AlibabaWdkWholesaleInboundorderCommitAPIResponse
func GetAlibabaWdkWholesaleInboundorderCommitAPIResponse() *AlibabaWdkWholesaleInboundorderCommitAPIResponse {
	return poolAlibabaWdkWholesaleInboundorderCommitAPIResponse.Get().(*AlibabaWdkWholesaleInboundorderCommitAPIResponse)
}

// ReleaseAlibabaWdkWholesaleInboundorderCommitAPIResponse 将 AlibabaWdkWholesaleInboundorderCommitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkWholesaleInboundorderCommitAPIResponse(v *AlibabaWdkWholesaleInboundorderCommitAPIResponse) {
	v.Reset()
	poolAlibabaWdkWholesaleInboundorderCommitAPIResponse.Put(v)
}
