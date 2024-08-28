package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordCreateKeywordBatch 关键词添加
// alibaba.scbp.ad.keyword.create.keyword.batch
//
// 关键词添加
func AlibabaScbpAdKeywordCreateKeywordBatch(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordCreateKeywordBatchAPIRequest, resp *scbp.AlibabaScbpAdKeywordCreateKeywordBatchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
