package jym

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymItemExternalGoodsBatchSynoffsale 同步下架接口
// alibaba.jym.item.external.goods.batch.synoffsale
//
// 同步下架接口
func AlibabaJymItemExternalGoodsBatchSynoffsale(ctx context.Context, clt *core.SDKClient, req *jym.AlibabaJymItemExternalGoodsBatchSynoffsaleAPIRequest, resp *jym.AlibabaJymItemExternalGoodsBatchSynoffsaleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
