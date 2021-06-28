package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取定向效果报表数据 APIResponse
taobao.simba.rpt.targetingtageffect.get

获取定向效果报表数据
*/
type TaobaoSimbaRptTargetingtageffectGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_rpt_targetingtageffect_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 效果数据
    
    Results   []RptEffectEntityDTO `json:"results,omitempty" xml:"