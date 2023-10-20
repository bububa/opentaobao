package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolDeleteactivityAPIResponse 删除商品池活动 API返回值
// alibaba.hm.marketing.itempool.deleteactivity
//
// 删除商品池活动
type AlibabaHmMarketingItempoolDeleteactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolDeleteactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolDeleteactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItempoolDeleteactivityAPIResponseModel).Reset()
}

// AlibabaHmMarketingItempoolDeleteactivityAPIResponseModel is 删除商品池活动 成功返回结果
type AlibabaHmMarketingItempoolDeleteactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_deleteactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolDeleteactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItempoolDeleteactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItempoolDeleteactivityAPIResponse)
	},
}

// GetAlibabaHmMarketingItempoolDeleteactivityAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItempoolDeleteactivityAPIResponse
func GetAlibabaHmMarketingItempoolDeleteactivityAPIResponse() *AlibabaHmMarketingItempoolDeleteactivityAPIResponse {
	return poolAlibabaHmMarketingItempoolDeleteactivityAPIResponse.Get().(*AlibabaHmMarketingItempoolDeleteactivityAPIResponse)
}

// ReleaseAlibabaHmMarketingItempoolDeleteactivityAPIResponse 将 AlibabaHmMarketingItempoolDeleteactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItempoolDeleteactivityAPIResponse(v *AlibabaHmMarketingItempoolDeleteactivityAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItempoolDeleteactivityAPIResponse.Put(v)
}
