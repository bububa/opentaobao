package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolCreateactivityAPIResponse 添加商品池活动 API返回值
// alibaba.hm.marketing.itempool.createactivity
//
// 添加商品池活动
type AlibabaHmMarketingItempoolCreateactivityAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingItempoolCreateactivityAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolCreateactivityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingItempoolCreateactivityAPIResponseModel).Reset()
}

// AlibabaHmMarketingItempoolCreateactivityAPIResponseModel is 添加商品池活动 成功返回结果
type AlibabaHmMarketingItempoolCreateactivityAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_itempool_createactivity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建活动返回结果
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingItempoolCreateactivityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaHmMarketingItempoolCreateactivityAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingItempoolCreateactivityAPIResponse)
	},
}

// GetAlibabaHmMarketingItempoolCreateactivityAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingItempoolCreateactivityAPIResponse
func GetAlibabaHmMarketingItempoolCreateactivityAPIResponse() *AlibabaHmMarketingItempoolCreateactivityAPIResponse {
	return poolAlibabaHmMarketingItempoolCreateactivityAPIResponse.Get().(*AlibabaHmMarketingItempoolCreateactivityAPIResponse)
}

// ReleaseAlibabaHmMarketingItempoolCreateactivityAPIResponse 将 AlibabaHmMarketingItempoolCreateactivityAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingItempoolCreateactivityAPIResponse(v *AlibabaHmMarketingItempoolCreateactivityAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingItempoolCreateactivityAPIResponse.Put(v)
}
