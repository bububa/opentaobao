package bill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// Taobaobillbookbillsget 查询虚拟账户明细数据(自研发商家专用)
// taobao.bill.book.bills.get
//
// 查询虚拟账户明细数据
func Taobaobillbookbillsget(clt *core.SDKClient, req *bill.TaobaobillbookbillsgetAPIRequest, session string) (*bill.TaobaobillbookbillsgetAPIResponse, error) {
	var resp bill.TaobaobillbookbillsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
