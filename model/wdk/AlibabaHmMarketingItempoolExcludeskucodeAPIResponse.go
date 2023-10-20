package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolExcludeskucodeAPIResponse 商品池排除商品【品类优惠使用】 API返回值
// alibaba.hm.marketing.itempool.excludeskucode
//
// 品类优惠新增排除池
type AlibabaHmMarketingItempoolExcludeskucodeAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolExcludeskucodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolExcludeskucodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItempoolExcludeskucodeAPIResponseModel).Reset()
}

// AlibabaHmMarketingItempoolExcludeskucodeAPIResponseModel is 商品池排除商品【品类优惠使用】 成功返回结果
type AlibabaHmMarketingItempoolExcludeskucodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_excludeskucode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品报名活动的返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolExcludeskucodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItempoolExcludeskucodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItempoolExcludeskucodeAPIResponse)
	},
}

// GetAlibabaHmMarketingItempoolExcludeskucodeAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItempoolExcludeskucodeAPIResponse
func GetAlibabaHmMarketingItempoolExcludeskucodeAPIResponse() *AlibabaHmMarketingItempoolExcludeskucodeAPIResponse {
	return poolAlibabaHmMarketingItempoolExcludeskucodeAPIResponse.Get().(*AlibabaHmMarketingItempoolExcludeskucodeAPIResponse)
}

// ReleaseAlibabaHmMarketingItempoolExcludeskucodeAPIResponse 将 AlibabaHmMarketingItempoolExcludeskucodeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItempoolExcludeskucodeAPIResponse(v *AlibabaHmMarketingItempoolExcludeskucodeAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItempoolExcludeskucodeAPIResponse.Put(v)
}
