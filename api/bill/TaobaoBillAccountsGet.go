package bill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// Taobaobillaccountsget 查询费用科目信息(限自研商家)
// taobao.bill.accounts.get
//
// 查询费用账户信息
func Taobaobillaccountsget(clt *core.SDKClient, req *bill.TaobaobillaccountsgetAPIRequest, session string) (*bill.TaobaobillaccountsgetAPIResponse, error) {
	var resp bill.TaobaobillaccountsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
