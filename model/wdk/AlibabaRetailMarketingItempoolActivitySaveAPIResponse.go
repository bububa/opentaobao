package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivitySaveAPIResponse 【同城零售】保存商品池活动 API返回值
// alibaba.retail.marketing.itempool.activity.save
//
// 同城零售商品池活动保存
type AlibabaRetailMarketingItempoolActivitySaveAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingItempoolActivitySaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItempoolActivitySaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingItempoolActivitySaveAPIResponseModel).Reset()
}

// AlibabaRetailMarketingItempoolActivitySaveAPIResponseModel is 【同城零售】保存商品池活动 成功返回结果
type AlibabaRetailMarketingItempoolActivitySaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itempool_activity_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItempoolActivitySaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingItempoolActivitySaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingItempoolActivitySaveAPIResponse)
	},
}

// GetAlibabaRetailMarketingItempoolActivitySaveAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingItempoolActivitySaveAPIResponse
func GetAlibabaRetailMarketingItempoolActivitySaveAPIResponse() *AlibabaRetailMarketingItempoolActivitySaveAPIResponse {
	return poolAlibabaRetailMarketingItempoolActivitySaveAPIResponse.Get().(*AlibabaRetailMarketingItempoolActivitySaveAPIResponse)
}

// ReleaseAlibabaRetailMarketingItempoolActivitySaveAPIResponse 将 AlibabaRetailMarketingItempoolActivitySaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingItempoolActivitySaveAPIResponse(v *AlibabaRetailMarketingItempoolActivitySaveAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingItempoolActivitySaveAPIResponse.Put(v)
}
