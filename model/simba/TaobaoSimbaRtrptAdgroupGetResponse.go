package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推广组实时报表数据 APIResponse
taobao.simba.rtrpt.adgroup.get

获取推广组实时报表数据
*/
type TaobaoSimbaRtrptAdgroupGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_rtrpt_adgroup_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 1111
    
    Results   []RtRptResultEntityDTO `json:"results,omitempty" xml:"