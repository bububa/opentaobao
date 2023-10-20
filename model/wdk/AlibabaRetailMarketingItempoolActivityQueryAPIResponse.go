package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingItempoolActivityQueryAPIResponse 查询商品池活动【同城零售】 API返回值
// alibaba.retail.marketing.itempool.activity.query
//
// 查询商品池活动【同城零售】
type AlibabaRetailMarketingItempoolActivityQueryAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingItempoolActivityQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItempoolActivityQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingItempoolActivityQueryAPIResponseModel).Reset()
}

// AlibabaRetailMarketingItempoolActivityQueryAPIResponseModel is 查询商品池活动【同城零售】 成功返回结果
type AlibabaRetailMarketingItempoolActivityQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itempool_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 错误码
	ErrNumber string `json:"err_number,omitempty" xml:"err_number,omitempty"`
	// 响应体
	Data *ItemPoolPromotionActivityDto `json:"data,omitempty" xml:"data,omitempty"`
	// 成功标识
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingItempoolActivityQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrMessage = ""
	m.ErrNumber = ""
	m.Data = nil
	m.Succeed = false
}

var poolAlibabaRetailMarketingItempoolActivityQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingItempoolActivityQueryAPIResponse)
	},
}

// GetAlibabaRetailMarketingItempoolActivityQueryAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingItempoolActivityQueryAPIResponse
func GetAlibabaRetailMarketingItempoolActivityQueryAPIResponse() *AlibabaRetailMarketingItempoolActivityQueryAPIResponse {
	return poolAlibabaRetailMarketingItempoolActivityQueryAPIResponse.Get().(*AlibabaRetailMarketingItempoolActivityQueryAPIResponse)
}

// ReleaseAlibabaRetailMarketingItempoolActivityQueryAPIResponse 将 AlibabaRetailMarketingItempoolActivityQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingItempoolActivityQueryAPIResponse(v *AlibabaRetailMarketingItempoolActivityQueryAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingItempoolActivityQueryAPIResponse.Put(v)
}
