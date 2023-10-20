package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuScrollQueryAPIResponse 门店商品批量游标方式查询接口 API返回值
// alibaba.wdk.sku.scroll.query
//
// 通过游标方式批量获取门店商品信息，包括商品条码，商品名称，价格，会员价等信息。
type AlibabaWdkSkuScrollQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuScrollQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuScrollQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuScrollQueryAPIResponseModel).Reset()
}

// AlibabaWdkSkuScrollQueryAPIResponseModel is 门店商品批量游标方式查询接口 成功返回结果
type AlibabaWdkSkuScrollQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_scroll_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ApiScrollPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuScrollQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuScrollQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuScrollQueryAPIResponse)
	},
}

// GetAlibabaWdkSkuScrollQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuScrollQueryAPIResponse
func GetAlibabaWdkSkuScrollQueryAPIResponse() *AlibabaWdkSkuScrollQueryAPIResponse {
	return poolAlibabaWdkSkuScrollQueryAPIResponse.Get().(*AlibabaWdkSkuScrollQueryAPIResponse)
}

// ReleaseAlibabaWdkSkuScrollQueryAPIResponse 将 AlibabaWdkSkuScrollQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuScrollQueryAPIResponse(v *AlibabaWdkSkuScrollQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuScrollQueryAPIResponse.Put(v)
}
