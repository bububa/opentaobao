package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMessageaccountMesssageReplyAPIRequest
消息号下行回复接口 API请求
taobao.messageaccount.messsage.reply

外部 isv 调用该进口来进行消息号消息的回复 */
type TaobaoMessageaccountMesssageReplyAPIRequest struct {
	model.Params
	// 入参
	_param0 *ReplyMessageDto
}

// New
