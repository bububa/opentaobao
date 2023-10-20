package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivityCreateAPIResponse 创建商品池活动【同城零售】 API返回值
// alibaba.retail.marketing.itempool.activity.create
//
// 同城零售商品池活动创建
type AlibabaRetailMarketingItempoolActivityCreateAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingItempoolActivityCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItempoolActivityCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingItempoolActivityCreateAPIResponseModel).Reset()
}

// AlibabaRetailMarketingItempoolActivityCreateAPIResponseModel is 创建商品池活动【同城零售】 成功返回结果
type AlibabaRetailMarketingItempoolActivityCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itempool_activity_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItempoolActivityCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingItempoolActivityCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingItempoolActivityCreateAPIResponse)
	},
}

// GetAlibabaRetailMarketingItempoolActivityCreateAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingItempoolActivityCreateAPIResponse
func GetAlibabaRetailMarketingItempoolActivityCreateAPIResponse() *AlibabaRetailMarketingItempoolActivityCreateAPIResponse {
	return poolAlibabaRetailMarketingItempoolActivityCreateAPIResponse.Get().(*AlibabaRetailMarketingItempoolActivityCreateAPIResponse)
}

// ReleaseAlibabaRetailMarketingItempoolActivityCreateAPIResponse 将 AlibabaRetailMarketingItempoolActivityCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingItempoolActivityCreateAPIResponse(v *AlibabaRetailMarketingItempoolActivityCreateAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingItempoolActivityCreateAPIResponse.Put(v)
}
