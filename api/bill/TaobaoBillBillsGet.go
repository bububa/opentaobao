package bill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// Taobaobillbillsget 查询账单明细数据(自研发商家专用)
// taobao.bill.bills.get
//
// 查询账单明细数据
func Taobaobillbillsget(clt *core.SDKClient, req *bill.TaobaobillbillsgetAPIRequest, session string) (*bill.TaobaobillbillsgetAPIResponse, error) {
	var resp bill.TaobaobillbillsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
