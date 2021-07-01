package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimChatlogsImportAPIRequest
openim单聊消息导入 API请求
taobao.openim.chatlogs.import

提供openim账号的聊天消息导入功能 */
type TaobaoOpenimChatlogsImportAPIRequest struct {
	model.Params
	// 消息序列
	_messages []TextMessage
}

// New
