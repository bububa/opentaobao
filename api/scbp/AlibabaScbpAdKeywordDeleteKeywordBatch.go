package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordDeleteKeywordBatch 删除关键词
// alibaba.scbp.ad.keyword.delete.keyword.batch
//
// 删除关键词
func AlibabaScbpAdKeywordDeleteKeywordBatch(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordDeleteKeywordBatchAPIRequest, resp *scbp.AlibabaScbpAdKeywordDeleteKeywordBatchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
