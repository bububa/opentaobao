package bill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// TaobaoTaeBillsGet tae查询账单明细
// taobao.tae.bills.get
//
// tae查询账单明细
func TaobaoTaeBillsGet(clt *core.SDKClient, req *bill.TaobaoTaeBillsGetAPIRequest, session string) (*bill.TaobaoTaeBillsGetAPIResponse, error) {
	var resp bill.TaobaoTaeBillsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
