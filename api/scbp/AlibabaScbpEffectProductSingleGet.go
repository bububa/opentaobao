package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpEffectProductSingleGet 单个产品的报表
// alibaba.scbp.effect.product.single.get
//
// 单个产品的报表
func AlibabaScbpEffectProductSingleGet(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpEffectProductSingleGetAPIRequest, resp *scbp.AlibabaScbpEffectProductSingleGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
