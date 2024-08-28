package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymItemExternalGoodsBatchOnsale 交易猫外部商家批量上架商品接口
// alibaba.jym.item.external.goods.batch.onsale
//
// 供外部B端商家接入，提交批量上架商品请求，返回批量上架任务结果
func AlibabaJymItemExternalGoodsBatchOnsale(ctx context.Context, clt *core.SDKClient, req *product.AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest, resp *product.AlibabaJymItemExternalGoodsBatchOnsaleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
