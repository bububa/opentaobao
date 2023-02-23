package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AlibabaJymItemExternalGoodsBatchDelete 交易猫外部商家批量删除商品接口
// alibaba.jym.item.external.goods.batch.delete
//
// 交易猫外部商家批量删除商品接口
func AlibabaJymItemExternalGoodsBatchDelete(clt *core.SDKClient, req *product.AlibabaJymItemExternalGoodsBatchDeleteAPIRequest, session string) (*product.AlibabaJymItemExternalGoodsBatchDeleteAPIResponse, error) {
	var resp product.AlibabaJymItemExternalGoodsBatchDeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
