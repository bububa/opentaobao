package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkWholesaleOrderCommitAPIResponse 盒马帮采购确认订单接口 API返回值
// alibaba.wdk.wholesale.order.commit
//
// 盒马帮采购确认订单接口
type AlibabaWdkWholesaleOrderCommitAPIResponse struct {
	model.CommonResponse
	AlibabaWdkWholesaleOrderCommitAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkWholesaleOrderCommitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkWholesaleOrderCommitAPIResponseModel).Reset()
}

// AlibabaWdkWholesaleOrderCommitAPIResponseModel is 盒马帮采购确认订单接口 成功返回结果
type AlibabaWdkWholesaleOrderCommitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_wholesale_order_commit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkWholesaleOrderCommitApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkWholesaleOrderCommitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkWholesaleOrderCommitAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkWholesaleOrderCommitAPIResponse)
	},
}

// GetAlibabaWdkWholesaleOrderCommitAPIResponse 从 sync.Pool 获取 AlibabaWdkWholesaleOrderCommitAPIResponse
func GetAlibabaWdkWholesaleOrderCommitAPIResponse() *AlibabaWdkWholesaleOrderCommitAPIResponse {
	return poolAlibabaWdkWholesaleOrderCommitAPIResponse.Get().(*AlibabaWdkWholesaleOrderCommitAPIResponse)
}

// ReleaseAlibabaWdkWholesaleOrderCommitAPIResponse 将 AlibabaWdkWholesaleOrderCommitAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkWholesaleOrderCommitAPIResponse(v *AlibabaWdkWholesaleOrderCommitAPIResponse) {
	v.Reset()
	poolAlibabaWdkWholesaleOrderCommitAPIResponse.Put(v)
}
