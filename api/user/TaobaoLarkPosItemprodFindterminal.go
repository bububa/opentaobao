package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoLarkPosItemprodFindterminal 终端配置支持
// taobao.lark.pos.itemprod.findterminal
//
// 终端配置支持,读取如果不存在则创建和远程的连接配置并返回
func TaobaoLarkPosItemprodFindterminal(ctx context.Context, clt *core.SDKClient, req *user.TaobaoLarkPosItemprodFindterminalAPIRequest, resp *user.TaobaoLarkPosItemprodFindterminalAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
