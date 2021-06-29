package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广组效果报表数据对象 APIResponse
taobao.simba.rpt.adgroupeffect.get

推广组效果报表数据对象
*/
type TaobaoSimbaRptAdgroupeffectGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaRptAdgroupeffectGetResponse
}

type TaobaoSimbaRptAdgroupeffectGetResponse struct {
    XMLName xml.Name `xml:"simba_rpt_adgroupeffect_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 推广组效果报表数据对象
    
    RptAdgroupEffectList   string `json:"rpt_adgroup_effect_list,omitempty" xml:"rpt_adgroup_effect_list,omitempty"`

    
}
