package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuCombineskuAddAPIResponse 组合商品新增接口 API返回值
// alibaba.wdk.sku.combinesku.add
//
// 组合商品新增接口
type AlibabaWdkSkuCombineskuAddAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuCombineskuAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuCombineskuAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuCombineskuAddAPIResponseModel).Reset()
}

// AlibabaWdkSkuCombineskuAddAPIResponseModel is 组合商品新增接口 成功返回结果
type AlibabaWdkSkuCombineskuAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_combinesku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabaWdkSkuCombineskuAddApiResults `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuCombineskuAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuCombineskuAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuCombineskuAddAPIResponse)
	},
}

// GetAlibabaWdkSkuCombineskuAddAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuCombineskuAddAPIResponse
func GetAlibabaWdkSkuCombineskuAddAPIResponse() *AlibabaWdkSkuCombineskuAddAPIResponse {
	return poolAlibabaWdkSkuCombineskuAddAPIResponse.Get().(*AlibabaWdkSkuCombineskuAddAPIResponse)
}

// ReleaseAlibabaWdkSkuCombineskuAddAPIResponse 将 AlibabaWdkSkuCombineskuAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuCombineskuAddAPIResponse(v *AlibabaWdkSkuCombineskuAddAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuCombineskuAddAPIResponse.Put(v)
}
