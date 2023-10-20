package bill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// Taobaotaebillget tae查询单笔账单明细
// taobao.tae.bill.get
//
// 查询单笔账单明细
func Taobaotaebillget(clt *core.SDKClient, req *bill.TaobaotaebillgetAPIRequest, session string) (*bill.TaobaotaebillgetAPIResponse, error) {
	var resp bill.TaobaotaebillgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
