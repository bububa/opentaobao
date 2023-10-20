package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse 删除单品买赠活动 API返回值
// alibaba.retail.marketing.buygift.activity.delete
//
// 同城零售单品特价活动删除
type AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingBuygiftActivityDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingBuygiftActivityDeleteAPIResponseModel).Reset()
}

// AlibabaRetailMarketingBuygiftActivityDeleteAPIResponseModel is 删除单品买赠活动 成功返回结果
type AlibabaRetailMarketingBuygiftActivityDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_activity_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftActivityDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingBuygiftActivityDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse)
	},
}

// GetAlibabaRetailMarketingBuygiftActivityDeleteAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse
func GetAlibabaRetailMarketingBuygiftActivityDeleteAPIResponse() *AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse {
	return poolAlibabaRetailMarketingBuygiftActivityDeleteAPIResponse.Get().(*AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse)
}

// ReleaseAlibabaRetailMarketingBuygiftActivityDeleteAPIResponse 将 AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftActivityDeleteAPIResponse(v *AlibabaRetailMarketingBuygiftActivityDeleteAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftActivityDeleteAPIResponse.Put(v)
}
