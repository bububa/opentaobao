package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuProductScoreGet 产品质量分查询
// alibaba.icbu.product.score.get
//
// 产品质量分查询
func AlibabaIcbuProductScoreGet(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuProductScoreGetAPIRequest, resp *icbu.AlibabaIcbuProductScoreGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
