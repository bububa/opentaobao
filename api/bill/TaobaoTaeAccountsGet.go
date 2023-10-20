package bill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// Taobaotaeaccountsget tae查询费用科目信息
// taobao.tae.accounts.get
//
// tae查询费用科目信息
func Taobaotaeaccountsget(clt *core.SDKClient, req *bill.TaobaotaeaccountsgetAPIRequest, session string) (*bill.TaobaotaeaccountsgetAPIResponse, error) {
	var resp bill.TaobaotaeaccountsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
