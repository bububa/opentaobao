package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyBillDailyQuery 账单日汇总接口
// alibaba.tcls.aelophy.bill.daily.query
//
// 账单日汇总接口
func AlibabaTclsAelophyBillDailyQuery(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyBillDailyQueryAPIRequest, resp *wdk.AlibabaTclsAelophyBillDailyQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
