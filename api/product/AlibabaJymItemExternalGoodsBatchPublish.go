package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymItemExternalGoodsBatchPublish 交易猫外部商家批量发布商品接口
// alibaba.jym.item.external.goods.batch.publish
//
// 供外部B端商家接入，提交批量发布商品请求，返回批量发布任务结果
func AlibabaJymItemExternalGoodsBatchPublish(ctx context.Context, clt *core.SDKClient, req *product.AlibabaJymItemExternalGoodsBatchPublishAPIRequest, resp *product.AlibabaJymItemExternalGoodsBatchPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
