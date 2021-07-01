package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcMessagesConfirmAPIRequest
确认消费消息的状态 API请求
taobao.tmc.messages.confirm

确认消费消息的状态 */
type TaobaoTmcMessagesConfirmAPIRequest struct {
	model.Params
	// 分组名称，不传代表默认分组
	_groupName string
	// 处理成功的消息ID列表 最大 200个ID
	_sMessageIds []int64
	// 处理失败的消息ID列表--已废弃，无需传此字段
	_fMessageIds []int64
}

// New
