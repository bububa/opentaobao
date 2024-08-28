package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// TaobaoOfnSelfRecycleAuth 自助回收鉴权
// taobao.ofn.self.recycle.auth
//
// 自助回收鉴权
func TaobaoOfnSelfRecycleAuth(ctx context.Context, clt *core.SDKClient, req *trade.TaobaoOfnSelfRecycleAuthAPIRequest, resp *trade.TaobaoOfnSelfRecycleAuthAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
