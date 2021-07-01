package bill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

/* TaobaoBillBillsGet
查询账单明细数据(自研发商家专用)
taobao.bill.bills.get

查询账单明细数据 */
func TaobaoBillBillsGet(clt *core.SDKClient, req *bill.TaobaoBillBillsGetAPIRequest, session string) (*bill.TaobaoBillBillsGetAPIResponse, error) {
	var resp bill.TaobaoBillBillsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
