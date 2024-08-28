package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDealerRequisitionorderClose 供应商/分销商关闭采购申请/经销采购单
// taobao.fenxiao.dealer.requisitionorder.close
//
// 供应商或分销商关闭采购申请/经销采购单
func TaobaoFenxiaoDealerRequisitionorderClose(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDealerRequisitionorderCloseAPIRequest, resp *fenxiao.TaobaoFenxiaoDealerRequisitionorderCloseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
