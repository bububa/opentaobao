package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpAdGroupDeleteForbiddenProduct 删除屏蔽品
// alibaba.scbp.ad.group.delete.forbidden.product
//
// 删除屏蔽品
func AlibabaScbpAdGroupDeleteForbiddenProduct(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupDeleteForbiddenProductAPIRequest, resp *scbp.AlibabaScbpAdGroupDeleteForbiddenProductAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
