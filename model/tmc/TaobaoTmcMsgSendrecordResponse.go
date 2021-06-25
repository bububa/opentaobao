package tmc

import (
    "github.com/bububa/opentaobao/model"
)

/* 
消息发送记录查询 APIResponse
taobao.tmc.msg.sendrecord

查询单条消息发送记录，只返回返回条数和时间。
*/
type TaobaoTmcMsgSendrecordAPIResponse struct {
    model.CommonResponse
    Response *TaobaoTmcMsgSendrecordResponse `json:"taobao_tmc_msg_sendrecord_response,omitempty"`
}

type TaobaoTmcMsgSendrecordResponse struct {

    // tmc发送次数
    TmcSendTimes   int64 `json:"tmc_send_times,omitempty"`

    // 淘宝发送测试
    TbSendTimes   int64 `json:"tb_send_times,omitempty"`

    // 淘宝发送时间
    TbSendList   string `json:"tb_send_list,omitempty"`

    // TMC发送时间
    TmcSendList   string `json:"tmc_send_list,omitempty"`

}
