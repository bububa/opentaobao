package tmallsc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建主动预约开通条件 APIResponse
tmall.servicecenter.reservecond.create

1、设置主动预约开通条件
*/
type TmallServicecenterReservecondCreateAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterReservecondCreateResponse `json:"tmall_servicecenter_reservecond_create_response,omitempty"`
}

type TmallServicecenterReservecondCreateResponse struct {

    // 是否成功
    MsgSuccess   bool `json:"msg_success,omitempty"`

    // 返回信息
    MsgInfo   string `json:"msg_info,omitempty"`

    // 返回code
    MsgCode   string `json:"msg_code,omitempty"`

}
