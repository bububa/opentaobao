package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广计划下的推广组报表效果数据查询(只有汇总数据，无分类类型) APIResponse
taobao.simba.rpt.campadgroupeffect.get

推广计划下的推广组报表效果数据查询(只有汇总数据，无分类类型)
*/
type TaobaoSimbaRptCampadgroupeffectGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaRptCampadgroupeffectGetResponse
}

type TaobaoSimbaRptCampadgroupeffectGetResponse struct {
    XMLName xml.Name `xml:"simba_rpt_campadgroupeffect_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 推广计划下推广组的效果数据列表
    
    RptCampadgroupEffectList   string `json:"rpt_campadgroup_effect_list,omitempty" xml:"rpt_campadgroup_effect_list,omitempty"`

    
}
