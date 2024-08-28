package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoRefundGet 查询采购单退款信息
// taobao.fenxiao.refund.get
//
// 分销商或供应商可以查询某子单的退款信息，以及下游订单的退款信息
func TaobaoFenxiaoRefundGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoRefundGetAPIRequest, resp *fenxiao.TaobaoFenxiaoRefundGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
