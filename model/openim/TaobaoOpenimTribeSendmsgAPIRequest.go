package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimTribeSendmsgAPIRequest
发送群消息 API请求
taobao.openim.tribe.sendmsg

发送群消息，目前支持发送4种类型的群消息，普通文本，图片，语音，自定义消息 */
type TaobaoOpenimTribeSendmsgAPIRequest struct {
	model.Params
	// 群消息发送者，只有该群的成员才可以发送群消息
	_user *User
	// 群id
	_tribeId int64
	// 发送群消息
	_msg *TribeMsg
}

// New
