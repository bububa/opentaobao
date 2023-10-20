package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkwholesaleordercommit 盒马帮采购确认订单接口
// alibaba.wdk.wholesale.order.commit
//
// 盒马帮采购确认订单接口
func Alibabawdkwholesaleordercommit(clt *core.SDKClient, req *wdk.AlibabawdkwholesaleordercommitAPIRequest, session string) (*wdk.AlibabawdkwholesaleordercommitAPIResponse, error) {
	var resp wdk.AlibabawdkwholesaleordercommitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
