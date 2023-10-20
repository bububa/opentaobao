package bill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// TaobaoTaeBookBillGet tae查询单笔虚拟账户明细
// taobao.tae.book.bill.get
//
// tae查询单笔虚拟账户明细
func TaobaoTaeBookBillGet(clt *core.SDKClient, req *bill.TaobaoTaeBookBillGetAPIRequest, resp *bill.TaobaoTaeBookBillGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
