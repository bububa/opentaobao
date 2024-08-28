package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymItemExternalGoodsBatchtaskQuery 交易猫外部商家查询商品批量任务接口
// alibaba.jym.item.external.goods.batchtask.query
//
// 供外部B端商家接入，请求查询商品批量任务，返回商品批量任务查询结果
func AlibabaJymItemExternalGoodsBatchtaskQuery(ctx context.Context, clt *core.SDKClient, req *product.AlibabaJymItemExternalGoodsBatchtaskQueryAPIRequest, resp *product.AlibabaJymItemExternalGoodsBatchtaskQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
