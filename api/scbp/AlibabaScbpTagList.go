package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTagList 查询所有分组
// alibaba.scbp.tag.list
//
// 查询所有分组
func AlibabaScbpTagList(clt *core.SDKClient, req *scbp.AlibabaScbpTagListAPIRequest, resp *scbp.AlibabaScbpTagListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
