package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索人群实时报表 APIResponse
taobao.simba.rtrpt.targetingtag.get

获取搜搜人群实时报表
*/
type TaobaoSimbaRtrptTargetingtagGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_rtrpt_targetingtag_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 111
    
    Results   []RtRptResultEntityDTO `json:"results,omitempty" xml:"