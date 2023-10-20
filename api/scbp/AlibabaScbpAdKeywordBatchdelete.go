package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordBatchdelete 外贸直通车批量删除关键词
// alibaba.scbp.ad.keyword.batchdelete
//
// 外贸直通车批量删除关键词
func AlibabaScbpAdKeywordBatchdelete(clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordBatchdeleteAPIRequest, resp *scbp.AlibabaScbpAdKeywordBatchdeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
