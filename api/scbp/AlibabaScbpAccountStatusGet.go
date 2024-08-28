package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAccountStatusGet 查询账户级别关键词推广状态
// alibaba.scbp.account.status.get
//
// 查询账户级别关键词推广状态
func AlibabaScbpAccountStatusGet(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAccountStatusGetAPIRequest, resp *scbp.AlibabaScbpAccountStatusGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
