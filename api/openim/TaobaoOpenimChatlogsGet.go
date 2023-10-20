package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimChatlogsGet openim聊天记录查询接口
// taobao.openim.chatlogs.get
//
// 查询openim账号聊天记录
func TaobaoOpenimChatlogsGet(clt *core.SDKClient, req *openim.TaobaoOpenimChatlogsGetAPIRequest, resp *openim.TaobaoOpenimChatlogsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
