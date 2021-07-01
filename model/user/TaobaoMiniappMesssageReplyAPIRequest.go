package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMiniappMesssageReplyAPIRequest
轻店铺下行回复接口 API请求
taobao.miniapp.messsage.reply

外部 isv 调用该进口来进行轻店铺消息的回复 */
type TaobaoMiniappMesssageReplyAPIRequest struct {
	model.Params
	// 入参
	_param *ReplyMessageDto
}

// New
