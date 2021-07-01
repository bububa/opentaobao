package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribelogsImportAPIRequest
openim群聊天记录导入 API请求
taobao.openim.tribelogs.import

openim群聊天记录导入 */
type TaobaoOpenimTribelogsImportAPIRequest struct {
	model.Params
	// 群号。必须为已存在的群，且群主属于本app
	_tribeId int64
	// 消息列表
	_messages []TribeTextMessage
}

// New
