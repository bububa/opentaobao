package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkSkuStoreskuScrollQueryAPIResponse 门店商品批量查询接口 API返回值
// alibaba.wdk.sku.storesku.scroll.query
//
// 提供门店商品批量查询接口
type AlibabaWdkSkuStoreskuScrollQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuStoreskuScrollQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkSkuStoreskuScrollQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkSkuStoreskuScrollQueryAPIResponseModel).Reset()
}

// AlibabaWdkSkuStoreskuScrollQueryAPIResponseModel is 门店商品批量查询接口 成功返回结果
type AlibabaWdkSkuStoreskuScrollQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_storesku_scroll_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *ApiScrollPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkSkuStoreskuScrollQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkSkuStoreskuScrollQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkSkuStoreskuScrollQueryAPIResponse)
	},
}

// GetAlibabaWdkSkuStoreskuScrollQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkSkuStoreskuScrollQueryAPIResponse
func GetAlibabaWdkSkuStoreskuScrollQueryAPIResponse() *AlibabaWdkSkuStoreskuScrollQueryAPIResponse {
	return poolAlibabaWdkSkuStoreskuScrollQueryAPIResponse.Get().(*AlibabaWdkSkuStoreskuScrollQueryAPIResponse)
}

// ReleaseAlibabaWdkSkuStoreskuScrollQueryAPIResponse 将 AlibabaWdkSkuStoreskuScrollQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkSkuStoreskuScrollQueryAPIResponse(v *AlibabaWdkSkuStoreskuScrollQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkSkuStoreskuScrollQueryAPIResponse.Put(v)
}
