package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkWholesaleInboundorderCommit 盒马帮退货信息回传接口
// alibaba.wdk.wholesale.inboundorder.commit
//
// 盒马帮退货信息回传接口
func AlibabaWdkWholesaleInboundorderCommit(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkWholesaleInboundorderCommitAPIRequest, resp *wdk.AlibabaWdkWholesaleInboundorderCommitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
