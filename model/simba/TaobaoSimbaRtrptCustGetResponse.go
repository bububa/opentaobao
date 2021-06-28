package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取账户实时报表数据 APIResponse
taobao.simba.rtrpt.cust.get

获取账户实时报表数据
*/
type TaobaoSimbaRtrptCustGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_rtrpt_cust_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 11
    
    Results   []RtRptResultEntityDTO `json:"results,omitempty" xml:"