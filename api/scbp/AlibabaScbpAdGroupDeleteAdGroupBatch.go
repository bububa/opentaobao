package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdGroupDeleteAdGroupBatch 删除推广单元
// alibaba.scbp.ad.group.delete.ad.group.batch
//
// 删除推广单元
func AlibabaScbpAdGroupDeleteAdGroupBatch(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupDeleteAdGroupBatchAPIRequest, resp *scbp.AlibabaScbpAdGroupDeleteAdGroupBatchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
