package tmc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTmcMessagesConsumeAPIRequest
消费多条消息 API请求
taobao.tmc.messages.consume

消费多条消息。消费时如果没有返回消息，建议做控制，不要一直调api，浪费应用的流量。如对程序做好优化，若没有消息则，sleep 100ms 等。 */
type TaobaoTmcMessagesConsumeAPIRequest struct {
	model.Params
	// 用户分组名称，不传表示消费默认分组，如果应用没有设置用户分组，传入分组名称将会返回错误
	_groupName string
	// 每次批量消费消息的条数，最小值：10；最大值：200
	_quantity int64
}

// New
