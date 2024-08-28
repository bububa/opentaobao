package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoRefundQuery 批量查询采购退款
// taobao.fenxiao.refund.query
//
// 供应商按查询条件批量查询代销采购退款
func TaobaoFenxiaoRefundQuery(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoRefundQueryAPIRequest, resp *fenxiao.TaobaoFenxiaoRefundQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
