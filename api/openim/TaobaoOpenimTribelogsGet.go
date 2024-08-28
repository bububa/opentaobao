package openim

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// TaobaoOpenimTribelogsGet openim 群聊天记录导出接口
// taobao.openim.tribelogs.get
//
// 获取openim账号的群聊天记录
func TaobaoOpenimTribelogsGet(ctx context.Context, clt *core.SDKClient, req *openim.TaobaoOpenimTribelogsGetAPIRequest, resp *openim.TaobaoOpenimTribelogsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
