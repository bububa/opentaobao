package tmc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
消息发送记录查询 APIResponse
taobao.tmc.msg.sendrecord

查询单条消息发送记录，只返回返回条数和时间。
*/
type TaobaoTmcMsgSendrecordAPIResponse struct {
    model.CommonResponse
    TaobaoTmcMsgSendrecordResponse
}

type TaobaoTmcMsgSendrecordResponse struct {
    XMLName xml.Name `xml:"tmc_msg_sendrecord_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // tmc发送次数
    
    TmcSendTimes   int64 `json:"tmc_send_times,omitempty" xml:"tmc_send_times,omitempty"`

    
    // 淘宝发送测试
    
    TbSendTimes   int64 `json:"tb_send_times,omitempty" xml:"tb_send_times,omitempty"`

    
    // 淘宝发送时间
    
    TbSendList   string `json:"tb_send_list,omitempty" xml:"tb_send_list,omitempty"`

    
    // TMC发送时间
    
    TmcSendList   string `json:"tmc_send_list,omitempty" xml:"tmc_send_list,omitempty"`

    
}
