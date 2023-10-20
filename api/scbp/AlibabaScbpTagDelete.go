package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTagDelete 删除关键词分组
// alibaba.scbp.tag.delete
//
// 删除关键词分组
func AlibabaScbpTagDelete(clt *core.SDKClient, req *scbp.AlibabaScbpTagDeleteAPIRequest, resp *scbp.AlibabaScbpTagDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
