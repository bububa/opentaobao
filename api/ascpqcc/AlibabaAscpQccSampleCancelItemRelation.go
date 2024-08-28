package ascpqcc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpqcc"
)

// AlibabaAscpQccSampleCancelItemRelation 魅力惠样品解除父子商品关系
// alibaba.ascp.qcc.sample.cancel.item.relation
//
// 品控中心魅力惠样品解除父子商品关系
func AlibabaAscpQccSampleCancelItemRelation(ctx context.Context, clt *core.SDKClient, req *ascpqcc.AlibabaAscpQccSampleCancelItemRelationAPIRequest, resp *ascpqcc.AlibabaAscpQccSampleCancelItemRelationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
