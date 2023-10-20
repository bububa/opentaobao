package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordDelete 外贸直通车删除关键词
// alibaba.scbp.ad.keyword.delete
//
// 外贸直通车删除关键词
func AlibabaScbpAdKeywordDelete(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordDeleteAPIRequest, resp *scbp.AlibabaScbpAdKeywordDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
