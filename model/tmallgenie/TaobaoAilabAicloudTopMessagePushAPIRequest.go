package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopMessagePushAPIRequest
天猫精灵消息中心广播推送消息接口 API请求
taobao.ailab.aicloud.top.message.push

天猫精灵运营平台消息能力开放广播接口，主要开放给b端用户，用户可调用接口进行广播推送，将消息推送到天猫精灵设备或者天猫精灵APP中。 */
type TaobaoAilabAicloudTopMessagePushAPIRequest struct {
	model.Params
	// 消息推送请求
	_messageBroadcastRequest *MessageBroadcastRequest
	// 当前用户信息
	_userInfoContext *OpsRequestUserInfoContext
}

// New
