package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlinkMessageHistoryActionAPIRequest
操作历史消息 API请求
alibaba.alink.message.history.action

阿里智能操作历史消息 */
type AlibabaAlinkMessageHistoryActionAPIRequest struct {
	model.Params
	// 消息id
	_index string
	// 删除：delete,已读：read
	_action string
}

// New
