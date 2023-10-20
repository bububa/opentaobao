package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivityDeleteAPIResponse 删除商品池活动【同城零售】 API返回值
// alibaba.retail.marketing.itempool.activity.delete
//
// 同城零售商品池活动删除
type AlibabaRetailMarketingItempoolActivityDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingItempoolActivityDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItempoolActivityDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingItempoolActivityDeleteAPIResponseModel).Reset()
}

// AlibabaRetailMarketingItempoolActivityDeleteAPIResponseModel is 删除商品池活动【同城零售】 成功返回结果
type AlibabaRetailMarketingItempoolActivityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itempool_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItempoolActivityDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingItempoolActivityDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingItempoolActivityDeleteAPIResponse)
	},
}

// GetAlibabaRetailMarketingItempoolActivityDeleteAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingItempoolActivityDeleteAPIResponse
func GetAlibabaRetailMarketingItempoolActivityDeleteAPIResponse() *AlibabaRetailMarketingItempoolActivityDeleteAPIResponse {
	return poolAlibabaRetailMarketingItempoolActivityDeleteAPIResponse.Get().(*AlibabaRetailMarketingItempoolActivityDeleteAPIResponse)
}

// ReleaseAlibabaRetailMarketingItempoolActivityDeleteAPIResponse 将 AlibabaRetailMarketingItempoolActivityDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingItempoolActivityDeleteAPIResponse(v *AlibabaRetailMarketingItempoolActivityDeleteAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingItempoolActivityDeleteAPIResponse.Put(v)
}
