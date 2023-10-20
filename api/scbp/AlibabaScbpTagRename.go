package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTagRename 重命名关键词分组
// alibaba.scbp.tag.rename
//
// 重命名关键词分组
func AlibabaScbpTagRename(clt *core.SDKClient, req *scbp.AlibabaScbpTagRenameAPIRequest, resp *scbp.AlibabaScbpTagRenameAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
