package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMessageaccountMesssageNormalSendAPIRequest
下行普通消息 API请求
taobao.messageaccount.messsage.normal.send

消息号下行单个普通消息 */
type TaobaoMessageaccountMesssageNormalSendAPIRequest struct {
	model.Params
	// 下行普通消息
	_param *NormalMessageDto
}

// New
