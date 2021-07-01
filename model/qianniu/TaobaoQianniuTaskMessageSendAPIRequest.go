package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQianniuTaskMessageSendAPIRequest
发送任务提醒消息 API请求
taobao.qianniu.task.message.send

如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。消息会以任务消息的形式发给客户端。 */
type TaobaoQianniuTaskMessageSendAPIRequest struct {
	model.Params
	// 任务ID。如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。
	_taskId int64
	// 任务元id，如果taskid不为空，则只发给task对应的单个接收人。如果taskid为空，则发给metadata_id对应的所有接收人。
	_metadataId int64
}

// New
