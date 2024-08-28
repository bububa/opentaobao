package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkDgVegasTljStop 淘宝客-推广者-淘礼金暂停发放
// taobao.tbk.dg.vegas.tlj.stop
//
// 淘宝客推广者可对已经创建的淘礼金暂停发放
func TaobaoTbkDgVegasTljStop(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkDgVegasTljStopAPIRequest, resp *tbk.TaobaoTbkDgVegasTljStopAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
