package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabatclsaelophybilldailyquery 账单日汇总接口
// alibaba.tcls.aelophy.bill.daily.query
//
// 账单日汇总接口
func Alibabatclsaelophybilldailyquery(clt *core.SDKClient, req *wdk.AlibabatclsaelophybilldailyqueryAPIRequest, session string) (*wdk.AlibabatclsaelophybilldailyqueryAPIResponse, error) {
	var resp wdk.AlibabatclsaelophybilldailyqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
