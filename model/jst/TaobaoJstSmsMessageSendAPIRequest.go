package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsMessageSendAPIRequest
聚石塔数据paas短信发送接口 API请求
taobao.jst.sms.message.send

聚石塔短信PAAS场景中，ISV通过该API帮商家发送短信给用户。 */
type TaobaoJstSmsMessageSendAPIRequest struct {
	model.Params
	// 短信发送请求
	_sendMessageRequest *SendMessageRequest
}

// New
