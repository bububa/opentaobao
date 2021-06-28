package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-国家标签ID获取 APIResponse
alibaba.scbp.target.ad.plan.country.id.get

定向推广-国家标签ID获取
*/
type AlibabaScbpTargetAdPlanCountryIdGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_target_ad_plan_country_id_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 地区列表
    
    RegionList   []RegionView `json:"region_list,omitempty" xml:"