package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkWholesaleOrderCommit 盒马帮采购确认订单接口
// alibaba.wdk.wholesale.order.commit
//
// 盒马帮采购确认订单接口
func AlibabaWdkWholesaleOrderCommit(clt *core.SDKClient, req *wdk.AlibabaWdkWholesaleOrderCommitAPIRequest, resp *wdk.AlibabaWdkWholesaleOrderCommitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
