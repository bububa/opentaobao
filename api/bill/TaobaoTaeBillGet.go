package bill

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bill"
)

// TaobaoTaeBillGet tae查询单笔账单明细
// taobao.tae.bill.get
//
// 查询单笔账单明细
func TaobaoTaeBillGet(clt *core.SDKClient, req *bill.TaobaoTaeBillGetAPIRequest, resp *bill.TaobaoTaeBillGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
