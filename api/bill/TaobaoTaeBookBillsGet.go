package bill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// Taobaotaebookbillsget tae查询虚拟账户明细数据
// taobao.tae.book.bills.get
//
// tae查询虚拟账户明细数据
func Taobaotaebookbillsget(clt *core.SDKClient, req *bill.TaobaotaebookbillsgetAPIRequest, session string) (*bill.TaobaotaebookbillsgetAPIResponse, error) {
	var resp bill.TaobaotaebookbillsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
