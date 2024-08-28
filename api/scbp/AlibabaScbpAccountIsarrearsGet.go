package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAccountIsarrearsGet 查询关键词推广账户是否欠款
// alibaba.scbp.account.isarrears.get
//
// 查询关键词推广账户是否欠款
func AlibabaScbpAccountIsarrearsGet(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAccountIsarrearsGetAPIRequest, resp *scbp.AlibabaScbpAccountIsarrearsGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
