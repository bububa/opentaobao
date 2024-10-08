package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoYphOrdersGet 批量查询一盘货采购单信息
// taobao.fenxiao.yph.orders.get
//
// 一盘货商家批量查询采购单信息
func TaobaoFenxiaoYphOrdersGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoYphOrdersGetAPIRequest, resp *fenxiao.TaobaoFenxiaoYphOrdersGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
