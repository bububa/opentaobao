package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoYphRefundGet 一盘货商家单个查询退款单信息
// taobao.fenxiao.yph.refund.get
//
// 一盘货商家单个查询退款单信息
func TaobaoFenxiaoYphRefundGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoYphRefundGetAPIRequest, resp *fenxiao.TaobaoFenxiaoYphRefundGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
