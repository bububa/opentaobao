package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdGroupCreateForbiddenProduct 创建屏蔽品
// alibaba.scbp.ad.group.create.forbidden.product
//
// 创建屏蔽品
func AlibabaScbpAdGroupCreateForbiddenProduct(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupCreateForbiddenProductAPIRequest, resp *scbp.AlibabaScbpAdGroupCreateForbiddenProductAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
