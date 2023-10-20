package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordDeleteKeywordBatch 删除关键词
// alibaba.scbp.ad.keyword.delete.keyword.batch
//
// 删除关键词
func AlibabaScbpAdKeywordDeleteKeywordBatch(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest, resp *scbp.AlibabaScbpAdKeywordDeleteKeywordBatchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
