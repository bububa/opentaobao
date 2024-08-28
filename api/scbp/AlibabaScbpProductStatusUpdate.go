package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpProductStatusUpdate 修改P4P产品推广状态
// alibaba.scbp.product.status.update
//
// 修改P4P产品推广状态
func AlibabaScbpProductStatusUpdate(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpProductStatusUpdateAPIRequest, resp *scbp.AlibabaScbpProductStatusUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
