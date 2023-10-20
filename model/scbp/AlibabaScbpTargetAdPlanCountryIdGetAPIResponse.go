package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptargetadplancountryidgetAPIResponse 定向推广-国家标签ID获取 API返回值
// alibaba.scbp.target.ad.plan.country.id.get
//
// 定向推广-国家标签ID获取
type AlibabascbptargetadplancountryidgetAPIResponse struct {
	model.CommonResponse
	AlibabascbptargetadplancountryidgetAPIResponseModel
}

// AlibabascbptargetadplancountryidgetAPIResponseModel is 定向推广-国家标签ID获取 成功返回结果
type AlibabascbptargetadplancountryidgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_country_id_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 地区列表
	RegionList []RegionView `json:"region_list,omitempty" xml:"region_list>region_view,omitempty"`
}
