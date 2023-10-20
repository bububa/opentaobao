package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolAddcategoryAPIResponse 增加商品池里面的类目 API返回值
// alibaba.hm.marketing.itempool.addcategory
//
// 增加商品池里面的类目
type AlibabaHmMarketingItempoolAddcategoryAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolAddcategoryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolAddcategoryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItempoolAddcategoryAPIResponseModel).Reset()
}

// AlibabaHmMarketingItempoolAddcategoryAPIResponseModel is 增加商品池里面的类目 成功返回结果
type AlibabaHmMarketingItempoolAddcategoryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_addcategory_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品报名活动的返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolAddcategoryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItempoolAddcategoryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItempoolAddcategoryAPIResponse)
	},
}

// GetAlibabaHmMarketingItempoolAddcategoryAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItempoolAddcategoryAPIResponse
func GetAlibabaHmMarketingItempoolAddcategoryAPIResponse() *AlibabaHmMarketingItempoolAddcategoryAPIResponse {
	return poolAlibabaHmMarketingItempoolAddcategoryAPIResponse.Get().(*AlibabaHmMarketingItempoolAddcategoryAPIResponse)
}

// ReleaseAlibabaHmMarketingItempoolAddcategoryAPIResponse 将 AlibabaHmMarketingItempoolAddcategoryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItempoolAddcategoryAPIResponse(v *AlibabaHmMarketingItempoolAddcategoryAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItempoolAddcategoryAPIResponse.Put(v)
}
