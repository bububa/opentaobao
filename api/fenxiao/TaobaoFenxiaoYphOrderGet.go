package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoYphOrderGet 一盘货商家单个查询采购单信息
// taobao.fenxiao.yph.order.get
//
// 一盘货商家单个查询采购单信息
func TaobaoFenxiaoYphOrderGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoYphOrderGetAPIRequest, resp *fenxiao.TaobaoFenxiaoYphOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
