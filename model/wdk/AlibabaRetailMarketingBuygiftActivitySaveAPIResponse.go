package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivitySaveAPIResponse 【同城零售】单品买赠活动保存 API返回值
// alibaba.retail.marketing.buygift.activity.save
//
// 同城零售单品买赠活动保存
type AlibabaRetailMarketingBuygiftActivitySaveAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingBuygiftActivitySaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftActivitySaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingBuygiftActivitySaveAPIResponseModel).Reset()
}

// AlibabaRetailMarketingBuygiftActivitySaveAPIResponseModel is 【同城零售】单品买赠活动保存 成功返回结果
type AlibabaRetailMarketingBuygiftActivitySaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_activity_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftActivitySaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingBuygiftActivitySaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingBuygiftActivitySaveAPIResponse)
	},
}

// GetAlibabaRetailMarketingBuygiftActivitySaveAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftActivitySaveAPIResponse
func GetAlibabaRetailMarketingBuygiftActivitySaveAPIResponse() *AlibabaRetailMarketingBuygiftActivitySaveAPIResponse {
	return poolAlibabaRetailMarketingBuygiftActivitySaveAPIResponse.Get().(*AlibabaRetailMarketingBuygiftActivitySaveAPIResponse)
}

// ReleaseAlibabaRetailMarketingBuygiftActivitySaveAPIResponse 将 AlibabaRetailMarketingBuygiftActivitySaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftActivitySaveAPIResponse(v *AlibabaRetailMarketingBuygiftActivitySaveAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftActivitySaveAPIResponse.Put(v)
}
