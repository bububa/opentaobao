package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDealerRequisitionorderGet 批量查询采购申请/经销采购单
// taobao.fenxiao.dealer.requisitionorder.get
//
// 批量查询采购申请/经销采购单，目前支持供应商和分销商查询
func TaobaoFenxiaoDealerRequisitionorderGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDealerRequisitionorderGetAPIRequest, resp *fenxiao.TaobaoFenxiaoDealerRequisitionorderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
