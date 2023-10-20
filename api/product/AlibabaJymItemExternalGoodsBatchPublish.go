package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymItemExternalGoodsBatchPublish 交易猫外部商家批量发布商品接口
// alibaba.jym.item.external.goods.batch.publish
//
// 供外部B端商家接入，提交批量发布商品请求，返回批量发布任务结果
func AlibabaJymItemExternalGoodsBatchPublish(clt *core.SDKClient, req *product.AlibabaJymItemExternalGoodsBatchPublishAPIRequest, session string) (*product.AlibabaJymItemExternalGoodsBatchPublishAPIResponse, error) {
	var resp product.AlibabaJymItemExternalGoodsBatchPublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
