package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDealerRequisitionorderCreate 创建经销采购申请
// taobao.fenxiao.dealer.requisitionorder.create
//
// 创建经销采购申请
func TaobaoFenxiaoDealerRequisitionorderCreate(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDealerRequisitionorderCreateAPIRequest, resp *fenxiao.TaobaoFenxiaoDealerRequisitionorderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
