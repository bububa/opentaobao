package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推广词实时报表数据 APIResponse
taobao.simba.rtrpt.bidword.get

获取推广词报表数据
*/
type TaobaoSimbaRtrptBidwordGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_rtrpt_bidword_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // bidword result
    
    Results   []RtRptResultEntityDTO `json:"results,omitempty" xml:"