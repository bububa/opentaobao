package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpReckeywordSearch 推荐词-词推词
// alibaba.scbp.reckeyword.search
//
// 推荐词-词推词
func AlibabaScbpReckeywordSearch(clt *core.SDKClient, req *scbp.AlibabaScbpReckeywordSearchAPIRequest, resp *scbp.AlibabaScbpReckeywordSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
