package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// TaobaoOpenAccountSearch open account数据搜索
// taobao.open.account.search
//
// open account数据搜索
func TaobaoOpenAccountSearch(ctx context.Context, clt *core.SDKClient, req *user.TaobaoOpenAccountSearchAPIRequest, resp *user.TaobaoOpenAccountSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
