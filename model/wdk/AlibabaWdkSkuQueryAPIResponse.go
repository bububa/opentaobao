package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuQueryAPIResponse 查询商品 API返回值
// alibaba.wdk.sku.query
//
// 查询商品
type AlibabaWdkSkuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuQueryAPIResponseModel).Reset()
}

// AlibabaWdkSkuQueryAPIResponseModel is 查询商品 成功返回结果
type AlibabaWdkSkuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabaWdkSkuQueryApiResults `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuQueryAPIResponse)
	},
}

// GetAlibabaWdkSkuQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuQueryAPIResponse
func GetAlibabaWdkSkuQueryAPIResponse() *AlibabaWdkSkuQueryAPIResponse {
	return poolAlibabaWdkSkuQueryAPIResponse.Get().(*AlibabaWdkSkuQueryAPIResponse)
}

// ReleaseAlibabaWdkSkuQueryAPIResponse 将 AlibabaWdkSkuQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuQueryAPIResponse(v *AlibabaWdkSkuQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuQueryAPIResponse.Put(v)
}
