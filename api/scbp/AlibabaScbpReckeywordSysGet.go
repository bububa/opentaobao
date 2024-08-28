package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpReckeywordSysGet 系统推荐
// alibaba.scbp.reckeyword.sys.get
//
// 查询系统推荐词
func AlibabaScbpReckeywordSysGet(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpReckeywordSysGetAPIRequest, resp *scbp.AlibabaScbpReckeywordSysGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
