package bill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// TaobaoTaeBookBillGet tae查询单笔虚拟账户明细
// taobao.tae.book.bill.get
//
// tae查询单笔虚拟账户明细
func TaobaoTaeBookBillGet(ctx context.Context, clt *core.SDKClient, req *bill.TaobaoTaeBookBillGetAPIRequest, resp *bill.TaobaoTaeBookBillGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
