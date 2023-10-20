package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymItemExternalGoodsBatchOffsale 交易猫外部商家批量下架商品接口
// alibaba.jym.item.external.goods.batch.offsale
//
// 供外部B端商家接入，提交批量下架商品请求，返回批量下架任务结果
func AlibabaJymItemExternalGoodsBatchOffsale(clt *core.SDKClient, req *product.AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest, resp *product.AlibabaJymItemExternalGoodsBatchOffsaleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
