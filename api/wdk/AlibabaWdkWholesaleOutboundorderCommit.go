package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkWholesaleOutboundorderCommit 盒马帮发货信息回传接口
// alibaba.wdk.wholesale.outboundorder.commit
//
// 盒马帮发货信息回传接口
func AlibabaWdkWholesaleOutboundorderCommit(clt *core.SDKClient, req *wdk.AlibabaWdkWholesaleOutboundorderCommitAPIRequest, resp *wdk.AlibabaWdkWholesaleOutboundorderCommitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
