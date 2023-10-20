package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdktradeordersuccesscreate 五道口终态订单创建
// alibaba.wdk.trade.order.success.create
//
// 五道口终态订单创建
func Alibabawdktradeordersuccesscreate(clt *core.SDKClient, req *wdk.AlibabawdktradeordersuccesscreateAPIRequest, session string) (*wdk.AlibabawdktradeordersuccesscreateAPIResponse, error) {
	var resp wdk.AlibabawdktradeordersuccesscreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
