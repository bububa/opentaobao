package bill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// TaobaoTaeBillsGet tae查询账单明细
// taobao.tae.bills.get
//
// tae查询账单明细
func TaobaoTaeBillsGet(clt *core.SDKClient, req *bill.TaobaoTaeBillsGetAPIRequest, resp *bill.TaobaoTaeBillsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
