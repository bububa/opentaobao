package bill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// TaobaoTaeBookBillsGet tae查询虚拟账户明细数据
// taobao.tae.book.bills.get
//
// tae查询虚拟账户明细数据
func TaobaoTaeBookBillsGet(ctx context.Context, clt *core.SDKClient, req *bill.TaobaoTaeBookBillsGetAPIRequest, resp *bill.TaobaoTaeBookBillsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
