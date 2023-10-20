package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribeSendmsg 发送群消息
// taobao.openim.tribe.sendmsg
//
// 发送群消息，目前支持发送4种类型的群消息，普通文本，图片，语音，自定义消息
func TaobaoOpenimTribeSendmsg(clt *core.SDKClient, req *openim.TaobaoOpenimTribeSendmsgAPIRequest, resp *openim.TaobaoOpenimTribeSendmsgAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
