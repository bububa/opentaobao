package tmc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcMsgSendrecordAPIResponse 消息发送记录查询 API返回值
// taobao.tmc.msg.sendrecord
//
// 查询单条消息发送记录，只返回返回条数和时间。
type TaobaoTmcMsgSendrecordAPIResponse struct {
	model.CommonResponse
	TaobaoTmcMsgSendrecordAPIResponseModel
}

// TaobaoTmcMsgSendrecordAPIResponseModel is 消息发送记录查询 成功返回结果
type TaobaoTmcMsgSendrecordAPIResponseModel struct {
	XMLName xml.Name `xml:"tmc_msg_sendrecord_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 淘宝发送时间
	TbSendList string `json:"tb_send_list,omitempty" xml:"tb_send_list,omitempty"`
	// TMC发送时间
	TmcSendList string `json:"tmc_send_list,omitempty" xml:"tmc_send_list,omitempty"`
	// tmc发送次数
	TmcSendTimes int64 `json:"tmc_send_times,omitempty" xml:"tmc_send_times,omitempty"`
	// 淘宝发送测试
	TbSendTimes int64 `json:"tb_send_times,omitempty" xml:"tb_send_times,omitempty"`
}
