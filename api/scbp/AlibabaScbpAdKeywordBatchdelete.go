package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordBatchdelete 外贸直通车批量删除关键词
// alibaba.scbp.ad.keyword.batchdelete
//
// 外贸直通车批量删除关键词
func AlibabaScbpAdKeywordBatchdelete(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordBatchdeleteAPIRequest, resp *scbp.AlibabaScbpAdKeywordBatchdeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
