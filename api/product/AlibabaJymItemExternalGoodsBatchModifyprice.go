package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymItemExternalGoodsBatchModifyprice 交易猫外部商家批量商品改价接口
// alibaba.jym.item.external.goods.batch.modifyprice
//
// 供外部B端商家接入，提交批量商品改价请求，返回批量改价任务结果
func AlibabaJymItemExternalGoodsBatchModifyprice(clt *core.SDKClient, req *product.AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest, session string) (*product.AlibabaJymItemExternalGoodsBatchModifypriceAPIResponse, error) {
	var resp product.AlibabaJymItemExternalGoodsBatchModifypriceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
