package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimAppChatlogsGet openim应用聊天记录查询
// taobao.openim.app.chatlogs.get
//
// 查询openim应用的聊天记录
func TaobaoOpenimAppChatlogsGet(clt *core.SDKClient, req *openim.TaobaoOpenimAppChatlogsGetAPIRequest, resp *openim.TaobaoOpenimAppChatlogsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
