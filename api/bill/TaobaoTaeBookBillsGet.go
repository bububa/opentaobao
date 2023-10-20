package bill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// TaobaoTaeBookBillsGet tae查询虚拟账户明细数据
// taobao.tae.book.bills.get
//
// tae查询虚拟账户明细数据
func TaobaoTaeBookBillsGet(clt *core.SDKClient, req *bill.TaobaoTaeBookBillsGetAPIRequest, resp *bill.TaobaoTaeBookBillsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
