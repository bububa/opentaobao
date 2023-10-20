package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOrderFinanceBillQuery 资金合规商家账单
// alibaba.wdk.order.finance.bill.query
//
// 拉取资金合规商家账单
func AlibabaWdkOrderFinanceBillQuery(clt *core.SDKClient, req *wdk.AlibabaWdkOrderFinanceBillQueryAPIRequest, resp *wdk.AlibabaWdkOrderFinanceBillQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
