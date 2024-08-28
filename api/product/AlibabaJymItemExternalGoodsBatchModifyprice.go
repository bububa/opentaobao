package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymItemExternalGoodsBatchModifyprice 交易猫外部商家批量商品改价接口
// alibaba.jym.item.external.goods.batch.modifyprice
//
// 供外部B端商家接入，提交批量商品改价请求，返回批量改价任务结果
func AlibabaJymItemExternalGoodsBatchModifyprice(ctx context.Context, clt *core.SDKClient, req *product.AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest, resp *product.AlibabaJymItemExternalGoodsBatchModifypriceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
