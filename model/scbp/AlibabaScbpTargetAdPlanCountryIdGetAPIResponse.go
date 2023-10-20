package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanCountryIdGetAPIResponse 定向推广-国家标签ID获取 API返回值
// alibaba.scbp.target.ad.plan.country.id.get
//
// 定向推广-国家标签ID获取
type AlibabaScbpTargetAdPlanCountryIdGetAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTargetAdPlanCountryIdGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanCountryIdGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpTargetAdPlanCountryIdGetAPIResponseModel).Reset()
}

// AlibabaScbpTargetAdPlanCountryIdGetAPIResponseModel is 定向推广-国家标签ID获取 成功返回结果
type AlibabaScbpTargetAdPlanCountryIdGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_country_id_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 地区列表
	RegionList []RegionView `json:"region_list,omitempty" xml:"region_list>region_view,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpTargetAdPlanCountryIdGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RegionList = m.RegionList[:0]
}

var poolAlibabaScbpTargetAdPlanCountryIdGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpTargetAdPlanCountryIdGetAPIResponse)
	},
}

// GetAlibabaScbpTargetAdPlanCountryIdGetAPIResponse 从 sync.Pool 获取 AlibabaScbpTargetAdPlanCountryIdGetAPIResponse
func GetAlibabaScbpTargetAdPlanCountryIdGetAPIResponse() *AlibabaScbpTargetAdPlanCountryIdGetAPIResponse {
	return poolAlibabaScbpTargetAdPlanCountryIdGetAPIResponse.Get().(*AlibabaScbpTargetAdPlanCountryIdGetAPIResponse)
}

// ReleaseAlibabaScbpTargetAdPlanCountryIdGetAPIResponse 将 AlibabaScbpTargetAdPlanCountryIdGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpTargetAdPlanCountryIdGetAPIResponse(v *AlibabaScbpTargetAdPlanCountryIdGetAPIResponse) {
	v.Reset()
	poolAlibabaScbpTargetAdPlanCountryIdGetAPIResponse.Put(v)
}
