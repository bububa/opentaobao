package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenimImmsgPushAPIRequest
openim标准消息发送 API请求
taobao.openim.immsg.push

服务端对openim用户发送标准消息，包括文字、语音、图片等。 */
type TaobaoOpenimImmsgPushAPIRequest struct {
	model.Params
	// openim消息结构体
	_immsg *ImMsg
}

// New
