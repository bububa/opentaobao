package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolAddcategoryAPIResponse 增加商品池里面的类目 API返回值
// alibaba.wdk.marketing.itempool.addcategory
//
// 增加商品池里面的类目
type AlibabaWdkMarketingItempoolAddcategoryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItempoolAddcategoryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolAddcategoryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItempoolAddcategoryAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItempoolAddcategoryAPIResponseModel is 增加商品池里面的类目 成功返回结果
type AlibabaWdkMarketingItempoolAddcategoryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_addcategory_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品报名活动的返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolAddcategoryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItempoolAddcategoryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItempoolAddcategoryAPIResponse)
	},
}

// GetAlibabaWdkMarketingItempoolAddcategoryAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItempoolAddcategoryAPIResponse
func GetAlibabaWdkMarketingItempoolAddcategoryAPIResponse() *AlibabaWdkMarketingItempoolAddcategoryAPIResponse {
	return poolAlibabaWdkMarketingItempoolAddcategoryAPIResponse.Get().(*AlibabaWdkMarketingItempoolAddcategoryAPIResponse)
}

// ReleaseAlibabaWdkMarketingItempoolAddcategoryAPIResponse 将 AlibabaWdkMarketingItempoolAddcategoryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolAddcategoryAPIResponse(v *AlibabaWdkMarketingItempoolAddcategoryAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolAddcategoryAPIResponse.Put(v)
}
