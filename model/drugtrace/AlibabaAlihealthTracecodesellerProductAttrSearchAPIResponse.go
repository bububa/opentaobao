package drugtrace

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse 根据商品id获取商品属性 API返回值
// alibaba.alihealth.tracecodeseller.product.attr.search
//
// 根据商品id获取商品属性
type AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponseModel).Reset()
}

// AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponseModel is 根据商品id获取商品属性 成功返回结果
type AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_product_attr_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse)
	},
}

// GetAlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse 从 sync.Pool 获取 AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse
func GetAlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse() *AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse {
	return poolAlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse.Get().(*AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse)
}

// ReleaseAlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse 将 AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse(v *AlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthTracecodesellerProductAttrSearchAPIResponse.Put(v)
}
