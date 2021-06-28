package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线上小票号获取 APIResponse
alibaba.mj.oc.online.ticketno.get

线上小票号获取
*/
type AlibabaMjOcOnlineTicketnoGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mj_oc_online_ticketno_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    Data   string `json:"data,omitempty" xml:"