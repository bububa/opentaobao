package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdGroupUpdateAdGroupBatch 修改推广单元
// alibaba.scbp.ad.group.update.ad.group.batch
//
// 修改推广单元
func AlibabaScbpAdGroupUpdateAdGroupBatch(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupUpdateAdGroupBatchAPIRequest, resp *scbp.AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
