package bill

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// TaobaoTaeBillsGet tae查询账单明细
// taobao.tae.bills.get
//
// tae查询账单明细
func TaobaoTaeBillsGet(ctx context.Context, clt *core.SDKClient, req *bill.TaobaoTaeBillsGetAPIRequest, resp *bill.TaobaoTaeBillsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
