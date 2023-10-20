package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolAdditemAPIResponse 增加商品池里面的商品 API返回值
// alibaba.hm.marketing.itempool.additem
//
// 增加商品池里面的商品
type AlibabaHmMarketingItempoolAdditemAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolAdditemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolAdditemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItempoolAdditemAPIResponseModel).Reset()
}

// AlibabaHmMarketingItempoolAdditemAPIResponseModel is 增加商品池里面的商品 成功返回结果
type AlibabaHmMarketingItempoolAdditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_additem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品报名活动的返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolAdditemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItempoolAdditemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItempoolAdditemAPIResponse)
	},
}

// GetAlibabaHmMarketingItempoolAdditemAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItempoolAdditemAPIResponse
func GetAlibabaHmMarketingItempoolAdditemAPIResponse() *AlibabaHmMarketingItempoolAdditemAPIResponse {
	return poolAlibabaHmMarketingItempoolAdditemAPIResponse.Get().(*AlibabaHmMarketingItempoolAdditemAPIResponse)
}

// ReleaseAlibabaHmMarketingItempoolAdditemAPIResponse 将 AlibabaHmMarketingItempoolAdditemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItempoolAdditemAPIResponse(v *AlibabaHmMarketingItempoolAdditemAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItempoolAdditemAPIResponse.Put(v)
}
