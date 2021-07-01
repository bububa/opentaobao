package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcMsgSendrecordAPIRequest
消息发送记录查询 API请求
taobao.tmc.msg.sendrecord

查询单条消息发送记录，只返回返回条数和时间。 */
type TaobaoTmcMsgSendrecordAPIRequest struct {
	model.Params
	// 消息分组名
	_groupName string
	// TOPIC名称
	_topicName string
	// 消息主键ID
	_dataId string
}

// New
