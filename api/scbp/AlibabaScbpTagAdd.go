package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTagAdd 创建关键词分组
// alibaba.scbp.tag.add
//
// 创建关键词分组
func AlibabaScbpTagAdd(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpTagAddAPIRequest, resp *scbp.AlibabaScbpTagAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
