package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolExcludeskucodeAPIResponse 商品池排除商品【品类优惠使用】 API返回值
// alibaba.wdk.marketing.itempool.excludeskucode
//
// 品类优惠新增排除池
type AlibabaWdkMarketingItempoolExcludeskucodeAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingItempoolExcludeskucodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolExcludeskucodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingItempoolExcludeskucodeAPIResponseModel).Reset()
}

// AlibabaWdkMarketingItempoolExcludeskucodeAPIResponseModel is 商品池排除商品【品类优惠使用】 成功返回结果
type AlibabaWdkMarketingItempoolExcludeskucodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_itempool_excludeskucode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品报名活动的返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingItempoolExcludeskucodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkMarketingItempoolExcludeskucodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingItempoolExcludeskucodeAPIResponse)
	},
}

// GetAlibabaWdkMarketingItempoolExcludeskucodeAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingItempoolExcludeskucodeAPIResponse
func GetAlibabaWdkMarketingItempoolExcludeskucodeAPIResponse() *AlibabaWdkMarketingItempoolExcludeskucodeAPIResponse {
	return poolAlibabaWdkMarketingItempoolExcludeskucodeAPIResponse.Get().(*AlibabaWdkMarketingItempoolExcludeskucodeAPIResponse)
}

// ReleaseAlibabaWdkMarketingItempoolExcludeskucodeAPIResponse 将 AlibabaWdkMarketingItempoolExcludeskucodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolExcludeskucodeAPIResponse(v *AlibabaWdkMarketingItempoolExcludeskucodeAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolExcludeskucodeAPIResponse.Put(v)
}
