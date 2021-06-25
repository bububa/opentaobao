package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
线上小票号获取 APIResponse
alibaba.mj.oc.online.ticketno.get

线上小票号获取
*/
type AlibabaMjOcOnlineTicketnoGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaMjOcOnlineTicketnoGetResponse `json:"alibaba_mj_oc_online_ticketno_get_response,omitempty"`
}

type AlibabaMjOcOnlineTicketnoGetResponse struct {

    // 是否成功
    Data   string `json:"data,omitempty"`

    // 错误信息
    ErrorsMsg   string `json:"errors_msg,omitempty"`

    // 错误码
    ErrorsCode   int64 `json:"errors_code,omitempty"`

    // 调用是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
