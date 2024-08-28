package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDealerRequisitionorderAgree 供应商/分销商通过采购申请/经销采购单申请
// taobao.fenxiao.dealer.requisitionorder.agree
//
// 供应商或分销商通过采购申请/经销采购单审核
func TaobaoFenxiaoDealerRequisitionorderAgree(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDealerRequisitionorderAgreeAPIRequest, resp *fenxiao.TaobaoFenxiaoDealerRequisitionorderAgreeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
