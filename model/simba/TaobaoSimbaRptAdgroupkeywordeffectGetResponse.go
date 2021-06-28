package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
推广组下的词效果报表数据查询(明细数据不分类型查询) APIResponse
taobao.simba.rpt.adgroupkeywordeffect.get

推广组下的词效果报表数据查询(明细数据不分类型查询)
*/
type TaobaoSimbaRptAdgroupkeywordeffectGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_rpt_adgroupkeywordeffect_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 词效果数据返回结果
    
    RptAdgroupkeywordEffectList   string `json:"rpt_adgroupkeyword_effect_list,omitempty" xml:"