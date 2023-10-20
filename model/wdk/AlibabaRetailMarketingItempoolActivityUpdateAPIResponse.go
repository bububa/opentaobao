package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivityUpdateAPIResponse 更新商品池活动【同城零售】 API返回值
// alibaba.retail.marketing.itempool.activity.update
//
// 同城零售商品池活动更新
type AlibabaRetailMarketingItempoolActivityUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingItempoolActivityUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItempoolActivityUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingItempoolActivityUpdateAPIResponseModel).Reset()
}

// AlibabaRetailMarketingItempoolActivityUpdateAPIResponseModel is 更新商品池活动【同城零售】 成功返回结果
type AlibabaRetailMarketingItempoolActivityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itempool_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItempoolActivityUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingItempoolActivityUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingItempoolActivityUpdateAPIResponse)
	},
}

// GetAlibabaRetailMarketingItempoolActivityUpdateAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingItempoolActivityUpdateAPIResponse
func GetAlibabaRetailMarketingItempoolActivityUpdateAPIResponse() *AlibabaRetailMarketingItempoolActivityUpdateAPIResponse {
	return poolAlibabaRetailMarketingItempoolActivityUpdateAPIResponse.Get().(*AlibabaRetailMarketingItempoolActivityUpdateAPIResponse)
}

// ReleaseAlibabaRetailMarketingItempoolActivityUpdateAPIResponse 将 AlibabaRetailMarketingItempoolActivityUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingItempoolActivityUpdateAPIResponse(v *AlibabaRetailMarketingItempoolActivityUpdateAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingItempoolActivityUpdateAPIResponse.Put(v)
}
