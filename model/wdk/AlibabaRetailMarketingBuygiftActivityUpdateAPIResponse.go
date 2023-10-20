package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse 更新单品买赠活动 API返回值
// alibaba.retail.marketing.buygift.activity.update
//
// 同城零售单品买赠活动更新
type AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingBuygiftActivityUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaRetailMarketingBuygiftActivityUpdateAPIResponseModel).Reset()
}

// AlibabaRetailMarketingBuygiftActivityUpdateAPIResponseModel is 更新单品买赠活动 成功返回结果
type AlibabaRetailMarketingBuygiftActivityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaRetailMarketingBuygiftActivityUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaRetailMarketingBuygiftActivityUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse)
	},
}

// GetAlibabaRetailMarketingBuygiftActivityUpdateAPIResponse 从 sync.Pool 获取 AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse
func GetAlibabaRetailMarketingBuygiftActivityUpdateAPIResponse() *AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse {
	return poolAlibabaRetailMarketingBuygiftActivityUpdateAPIResponse.Get().(*AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse)
}

// ReleaseAlibabaRetailMarketingBuygiftActivityUpdateAPIResponse 将 AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaRetailMarketingBuygiftActivityUpdateAPIResponse(v *AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse) {
	v.Reset()
	poolAlibabaRetailMarketingBuygiftActivityUpdateAPIResponse.Put(v)
}
