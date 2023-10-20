package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdktraderefundsuccesscreate 五道口终态逆向订单创建
// alibaba.wdk.trade.refund.success.create
//
// 五道口终态逆向订单创建
func Alibabawdktraderefundsuccesscreate(clt *core.SDKClient, req *wdk.AlibabawdktraderefundsuccesscreateAPIRequest, session string) (*wdk.AlibabawdktraderefundsuccesscreateAPIResponse, error) {
	var resp wdk.AlibabawdktraderefundsuccesscreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
