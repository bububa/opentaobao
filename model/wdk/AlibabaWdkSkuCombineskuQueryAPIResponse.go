package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCombineskuQueryAPIResponse 组合商品查询接口 API返回值
// alibaba.wdk.sku.combinesku.query
//
// 查询组合商品接口
type AlibabaWdkSkuCombineskuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuCombineskuQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuCombineskuQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuCombineskuQueryAPIResponseModel).Reset()
}

// AlibabaWdkSkuCombineskuQueryAPIResponseModel is 组合商品查询接口 成功返回结果
type AlibabaWdkSkuCombineskuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_combinesku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabaWdkSkuCombineskuQueryApiResults `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuCombineskuQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuCombineskuQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuCombineskuQueryAPIResponse)
	},
}

// GetAlibabaWdkSkuCombineskuQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuCombineskuQueryAPIResponse
func GetAlibabaWdkSkuCombineskuQueryAPIResponse() *AlibabaWdkSkuCombineskuQueryAPIResponse {
	return poolAlibabaWdkSkuCombineskuQueryAPIResponse.Get().(*AlibabaWdkSkuCombineskuQueryAPIResponse)
}

// ReleaseAlibabaWdkSkuCombineskuQueryAPIResponse 将 AlibabaWdkSkuCombineskuQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuCombineskuQueryAPIResponse(v *AlibabaWdkSkuCombineskuQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuCombineskuQueryAPIResponse.Put(v)
}
