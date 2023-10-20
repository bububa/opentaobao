package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolStairQueryitemAPIResponse 换购商品查询 API返回值
// alibaba.hm.marketing.itempool.stair.queryitem
//
// 换购商品查询
type AlibabaHmMarketingItempoolStairQueryitemAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolStairQueryitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolStairQueryitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItempoolStairQueryitemAPIResponseModel).Reset()
}

// AlibabaHmMarketingItempoolStairQueryitemAPIResponseModel is 换购商品查询 成功返回结果
type AlibabaHmMarketingItempoolStairQueryitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_stair_queryitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Result *MarketPageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolStairQueryitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItempoolStairQueryitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItempoolStairQueryitemAPIResponse)
	},
}

// GetAlibabaHmMarketingItempoolStairQueryitemAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItempoolStairQueryitemAPIResponse
func GetAlibabaHmMarketingItempoolStairQueryitemAPIResponse() *AlibabaHmMarketingItempoolStairQueryitemAPIResponse {
	return poolAlibabaHmMarketingItempoolStairQueryitemAPIResponse.Get().(*AlibabaHmMarketingItempoolStairQueryitemAPIResponse)
}

// ReleaseAlibabaHmMarketingItempoolStairQueryitemAPIResponse 将 AlibabaHmMarketingItempoolStairQueryitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItempoolStairQueryitemAPIResponse(v *AlibabaHmMarketingItempoolStairQueryitemAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItempoolStairQueryitemAPIResponse.Put(v)
}
