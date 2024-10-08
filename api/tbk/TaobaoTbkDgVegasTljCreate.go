package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgVegasTljCreate 淘宝客-推广者-淘礼金创建
// taobao.tbk.dg.vegas.tlj.create
//
// 创建淘礼金
func TaobaoTbkDgVegasTljCreate(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkDgVegasTljCreateAPIRequest, resp *tbk.TaobaoTbkDgVegasTljCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
