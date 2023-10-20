package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolStairRemoveitemAPIResponse 删除换购活动商品 API返回值
// alibaba.hm.marketing.itempool.stair.removeitem
//
// 删除换购商品
type AlibabaHmMarketingItempoolStairRemoveitemAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolStairRemoveitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolStairRemoveitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItempoolStairRemoveitemAPIResponseModel).Reset()
}

// AlibabaHmMarketingItempoolStairRemoveitemAPIResponseModel is 删除换购活动商品 成功返回结果
type AlibabaHmMarketingItempoolStairRemoveitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_stair_removeitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolStairRemoveitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItempoolStairRemoveitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItempoolStairRemoveitemAPIResponse)
	},
}

// GetAlibabaHmMarketingItempoolStairRemoveitemAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItempoolStairRemoveitemAPIResponse
func GetAlibabaHmMarketingItempoolStairRemoveitemAPIResponse() *AlibabaHmMarketingItempoolStairRemoveitemAPIResponse {
	return poolAlibabaHmMarketingItempoolStairRemoveitemAPIResponse.Get().(*AlibabaHmMarketingItempoolStairRemoveitemAPIResponse)
}

// ReleaseAlibabaHmMarketingItempoolStairRemoveitemAPIResponse 将 AlibabaHmMarketingItempoolStairRemoveitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItempoolStairRemoveitemAPIResponse(v *AlibabaHmMarketingItempoolStairRemoveitemAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItempoolStairRemoveitemAPIResponse.Put(v)
}
