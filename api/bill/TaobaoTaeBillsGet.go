package bill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// Taobaotaebillsget tae查询账单明细
// taobao.tae.bills.get
//
// tae查询账单明细
func Taobaotaebillsget(clt *core.SDKClient, req *bill.TaobaotaebillsgetAPIRequest, session string) (*bill.TaobaotaebillsgetAPIResponse, error) {
	var resp bill.TaobaotaebillsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
