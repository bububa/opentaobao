package bill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// Taobaotaebookbillget tae查询单笔虚拟账户明细
// taobao.tae.book.bill.get
//
// tae查询单笔虚拟账户明细
func Taobaotaebookbillget(clt *core.SDKClient, req *bill.TaobaotaebookbillgetAPIRequest, session string) (*bill.TaobaotaebookbillgetAPIResponse, error) {
	var resp bill.TaobaotaebookbillgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
