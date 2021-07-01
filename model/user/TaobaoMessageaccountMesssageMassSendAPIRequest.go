package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMessageaccountMesssageMassSendAPIRequest
消息号开放-消息群发 API请求
taobao.messageaccount.messsage.mass.send

外部 isv 调用该进口来进行消息号消息的群发 */
type TaobaoMessageaccountMesssageMassSendAPIRequest struct {
	model.Params
	// 参数
	_param *MassMessageDto
}

// New
