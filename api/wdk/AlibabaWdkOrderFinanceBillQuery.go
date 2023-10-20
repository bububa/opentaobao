package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkorderfinancebillquery 资金合规商家账单
// alibaba.wdk.order.finance.bill.query
//
// 拉取资金合规商家账单
func Alibabawdkorderfinancebillquery(clt *core.SDKClient, req *wdk.AlibabawdkorderfinancebillqueryAPIRequest, session string) (*wdk.AlibabawdkorderfinancebillqueryAPIResponse, error) {
	var resp wdk.AlibabawdkorderfinancebillqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
