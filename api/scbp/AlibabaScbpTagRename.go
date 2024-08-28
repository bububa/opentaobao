package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTagRename 重命名关键词分组
// alibaba.scbp.tag.rename
//
// 重命名关键词分组
func AlibabaScbpTagRename(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpTagRenameAPIRequest, resp *scbp.AlibabaScbpTagRenameAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
