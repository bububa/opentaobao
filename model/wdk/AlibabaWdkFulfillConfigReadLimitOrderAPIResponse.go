package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkFulfillConfigReadLimitOrderAPIResponse 根据仓code查询仓限单配置 API返回值
// alibaba.wdk.fulfill.config.read.limit.order
//
// 根据仓code查询仓限单配置
type AlibabaWdkFulfillConfigReadLimitOrderAPIResponse struct {
	model.CommonResponse
	AlibabaWdkFulfillConfigReadLimitOrderAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillConfigReadLimitOrderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkFulfillConfigReadLimitOrderAPIResponseModel).Reset()
}

// AlibabaWdkFulfillConfigReadLimitOrderAPIResponseModel is 根据仓code查询仓限单配置 成功返回结果
type AlibabaWdkFulfillConfigReadLimitOrderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_fulfill_config_read_limit_order_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkFulfillConfigReadLimitOrderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolAlibabaWdkFulfillConfigReadLimitOrderAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkFulfillConfigReadLimitOrderAPIResponse)
	},
}

// GetAlibabaWdkFulfillConfigReadLimitOrderAPIResponse 从 sync.Pool 获取 AlibabaWdkFulfillConfigReadLimitOrderAPIResponse
func GetAlibabaWdkFulfillConfigReadLimitOrderAPIResponse() *AlibabaWdkFulfillConfigReadLimitOrderAPIResponse {
	return poolAlibabaWdkFulfillConfigReadLimitOrderAPIResponse.Get().(*AlibabaWdkFulfillConfigReadLimitOrderAPIResponse)
}

// ReleaseAlibabaWdkFulfillConfigReadLimitOrderAPIResponse 将 AlibabaWdkFulfillConfigReadLimitOrderAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkFulfillConfigReadLimitOrderAPIResponse(v *AlibabaWdkFulfillConfigReadLimitOrderAPIResponse) {
	v.Reset()
	poolAlibabaWdkFulfillConfigReadLimitOrderAPIResponse.Put(v)
}
