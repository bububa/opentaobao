package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdKeywordUpdateKeywordStatusBatch 修改关键词状态
// alibaba.scbp.ad.keyword.update.keyword.status.batch
//
// 修改关键词状态
func AlibabaScbpAdKeywordUpdateKeywordStatusBatch(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIRequest, resp *scbp.AlibabaScbpAdKeywordUpdateKeywordStatusBatchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
