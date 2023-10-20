package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkBmStockPublishAPIResponse 品牌营销涉及到的商品的库存同步接口 API返回值
// alibaba.wdk.bm.stock.publish
//
// 用于操作sku的库存
type AlibabaWdkBmStockPublishAPIResponse struct {
	model.CommonResponse
	AlibabaWdkBmStockPublishAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkBmStockPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkBmStockPublishAPIResponseModel).Reset()
}

// AlibabaWdkBmStockPublishAPIResponseModel is 品牌营销涉及到的商品的库存同步接口 成功返回结果
type AlibabaWdkBmStockPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_bm_stock_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *BmResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkBmStockPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkBmStockPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkBmStockPublishAPIResponse)
	},
}

// GetAlibabaWdkBmStockPublishAPIResponse 从 sync.Pool 获取 AlibabaWdkBmStockPublishAPIResponse
func GetAlibabaWdkBmStockPublishAPIResponse() *AlibabaWdkBmStockPublishAPIResponse {
	return poolAlibabaWdkBmStockPublishAPIResponse.Get().(*AlibabaWdkBmStockPublishAPIResponse)
}

// ReleaseAlibabaWdkBmStockPublishAPIResponse 将 AlibabaWdkBmStockPublishAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkBmStockPublishAPIResponse(v *AlibabaWdkBmStockPublishAPIResponse) {
	v.Reset()
	poolAlibabaWdkBmStockPublishAPIResponse.Put(v)
}
