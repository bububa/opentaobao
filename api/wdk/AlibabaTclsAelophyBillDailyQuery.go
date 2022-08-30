package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaTclsAelophyBillDailyQuery 账单日汇总接口
// alibaba.tcls.aelophy.bill.daily.query
//
// 账单日汇总接口
func AlibabaTclsAelophyBillDailyQuery(clt *core.SDKClient, req *wdk.AlibabaTclsAelophyBillDailyQueryAPIRequest, session string) (*wdk.AlibabaTclsAelophyBillDailyQueryAPIResponse, error) {
	var resp wdk.AlibabaTclsAelophyBillDailyQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
