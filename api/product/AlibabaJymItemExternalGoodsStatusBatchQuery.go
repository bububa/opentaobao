package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymItemExternalGoodsStatusBatchQuery 交易猫外部商家商品状态批量查询接口
// alibaba.jym.item.external.goods.status.batch.query
//
// 供外部B端商家接入，请求查询商品状态，返回商品状态查询结果
func AlibabaJymItemExternalGoodsStatusBatchQuery(ctx context.Context, clt *core.SDKClient, req *product.AlibabaJymItemExternalGoodsStatusBatchQueryAPIRequest, resp *product.AlibabaJymItemExternalGoodsStatusBatchQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
