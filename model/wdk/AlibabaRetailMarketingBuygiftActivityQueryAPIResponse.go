package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivityQueryAPIResponse 查询单品买赠活动【同城零售】 API返回值
// alibaba.retail.marketing.buygift.activity.query
//
// 查询单品买赠活动【同城零售】
type AlibabaRetailMarketingBuygiftActivityQueryAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingBuygiftActivityQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftActivityQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingBuygiftActivityQueryAPIResponseModel).Reset()
}

// AlibabaRetailMarketingBuygiftActivityQueryAPIResponseModel is 查询单品买赠活动【同城零售】 成功返回结果
type AlibabaRetailMarketingBuygiftActivityQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 错误码
	ErrNumber string `json:"err_number,omitempty" xml:"err_number,omitempty"`
	// 买赠活动查询结果
	Data *BuyGiftPromotionActivityDto `json:"data,omitempty" xml:"data,omitempty"`
	// 成功标识
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftActivityQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrMessage = ""
	m.ErrNumber = ""
	m.Data = nil
	m.Succeed = false
}

var poolAlibabaRetailMarketingBuygiftActivityQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingBuygiftActivityQueryAPIResponse)
	},
}

// GetAlibabaRetailMarketingBuygiftActivityQueryAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftActivityQueryAPIResponse
func GetAlibabaRetailMarketingBuygiftActivityQueryAPIResponse() *AlibabaRetailMarketingBuygiftActivityQueryAPIResponse {
	return poolAlibabaRetailMarketingBuygiftActivityQueryAPIResponse.Get().(*AlibabaRetailMarketingBuygiftActivityQueryAPIResponse)
}

// ReleaseAlibabaRetailMarketingBuygiftActivityQueryAPIResponse 将 AlibabaRetailMarketingBuygiftActivityQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftActivityQueryAPIResponse(v *AlibabaRetailMarketingBuygiftActivityQueryAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftActivityQueryAPIResponse.Put(v)
}
