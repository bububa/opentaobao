package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkItemInfoGet 淘宝客-公用-淘宝客商品详情查询(简版)
// taobao.tbk.item.info.get
//
// 淘宝客商品详情查询（简版）
func TaobaoTbkItemInfoGet(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkItemInfoGetAPIRequest, resp *tbk.TaobaoTbkItemInfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
