package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-国家标签ID获取 APIResponse
alibaba.scbp.target.ad.plan.country.id.get

定向推广-国家标签ID获取
*/
type AlibabaScbpTargetAdPlanCountryIdGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpTargetAdPlanCountryIdGetResponse `json:"alibaba_scbp_target_ad_plan_country_id_get_response,omitempty"`
}

type AlibabaScbpTargetAdPlanCountryIdGetResponse struct {

    // 地区列表
    RegionList   []RegionView `json:"region_list,omitempty"`

}
