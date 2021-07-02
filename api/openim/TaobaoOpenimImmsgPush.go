package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimImmsgPush openim标准消息发送
// taobao.openim.immsg.push
//
// 服务端对openim用户发送标准消息，包括文字、语音、图片等。
func TaobaoOpenimImmsgPush(clt *core.SDKClient, req *openim.TaobaoOpenimImmsgPushAPIRequest, session string) (*openim.TaobaoOpenimImmsgPushAPIResponse, error) {
	var resp openim.TaobaoOpenimImmsgPushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
