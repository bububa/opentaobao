package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymItemExternalGoodsDetailQuery 交易猫外部商家商品详情查询接口
// alibaba.jym.item.external.goods.detail.query
//
// 供外部B端商家接入，请求查询商品详情，返回商品详情查询结果
func AlibabaJymItemExternalGoodsDetailQuery(ctx context.Context, clt *core.SDKClient, req *product.AlibabaJymItemExternalGoodsDetailQueryAPIRequest, resp *product.AlibabaJymItemExternalGoodsDetailQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
