package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuMerchantskuScrollQueryAPIResponse 商家商品批量查询接口 API返回值
// alibaba.wdk.sku.merchantsku.scroll.query
//
// 提供主档商品数据接口查询
type AlibabaWdkSkuMerchantskuScrollQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuMerchantskuScrollQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuMerchantskuScrollQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuMerchantskuScrollQueryAPIResponseModel).Reset()
}

// AlibabaWdkSkuMerchantskuScrollQueryAPIResponseModel is 商家商品批量查询接口 成功返回结果
type AlibabaWdkSkuMerchantskuScrollQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_merchantsku_scroll_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果对象
	Result *ApiScrollPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuMerchantskuScrollQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuMerchantskuScrollQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuMerchantskuScrollQueryAPIResponse)
	},
}

// GetAlibabaWdkSkuMerchantskuScrollQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuMerchantskuScrollQueryAPIResponse
func GetAlibabaWdkSkuMerchantskuScrollQueryAPIResponse() *AlibabaWdkSkuMerchantskuScrollQueryAPIResponse {
	return poolAlibabaWdkSkuMerchantskuScrollQueryAPIResponse.Get().(*AlibabaWdkSkuMerchantskuScrollQueryAPIResponse)
}

// ReleaseAlibabaWdkSkuMerchantskuScrollQueryAPIResponse 将 AlibabaWdkSkuMerchantskuScrollQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuMerchantskuScrollQueryAPIResponse(v *AlibabaWdkSkuMerchantskuScrollQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuMerchantskuScrollQueryAPIResponse.Put(v)
}
