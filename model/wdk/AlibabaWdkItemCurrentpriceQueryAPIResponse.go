package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemCurrentpriceQueryAPIResponse 查询商品当前价格 API返回值
// alibaba.wdk.item.currentprice.query
//
// 通过渠道店id/sku编码/渠道查询商品当前价格，一次请求商品数量&lt;=20,返回结果key为skuCode
type AlibabaWdkItemCurrentpriceQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemCurrentpriceQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkItemCurrentpriceQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkItemCurrentpriceQueryAPIResponseModel).Reset()
}

// AlibabaWdkItemCurrentpriceQueryAPIResponseModel is 查询商品当前价格 成功返回结果
type AlibabaWdkItemCurrentpriceQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_currentprice_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaWdkItemCurrentpriceQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkItemCurrentpriceQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkItemCurrentpriceQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkItemCurrentpriceQueryAPIResponse)
	},
}

// GetAlibabaWdkItemCurrentpriceQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkItemCurrentpriceQueryAPIResponse
func GetAlibabaWdkItemCurrentpriceQueryAPIResponse() *AlibabaWdkItemCurrentpriceQueryAPIResponse {
	return poolAlibabaWdkItemCurrentpriceQueryAPIResponse.Get().(*AlibabaWdkItemCurrentpriceQueryAPIResponse)
}

// ReleaseAlibabaWdkItemCurrentpriceQueryAPIResponse 将 AlibabaWdkItemCurrentpriceQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkItemCurrentpriceQueryAPIResponse(v *AlibabaWdkItemCurrentpriceQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkItemCurrentpriceQueryAPIResponse.Put(v)
}
