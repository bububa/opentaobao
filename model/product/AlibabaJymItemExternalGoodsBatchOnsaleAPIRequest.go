package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest 交易猫外部商家批量上架商品接口 API请求
// alibaba.jym.item.external.goods.batch.onsale
//
// 供外部B端商家接入，提交批量上架商品请求，返回批量上架任务结果
type AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest struct {
	model.Params
	// 商品批量上架接口入参
	_goodsOnSaleCommand *GoodsOnSaleCommandDto
}

// NewAlibabaJymItemExternalGoodsBatchOnsaleRequest 初始化AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest对象
func NewAlibabaJymItemExternalGoodsBatchOnsaleRequest() *AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest {
	return &AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batch.onsale"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetGoodsOnSaleCommand is GoodsOnSaleCommand Setter
// 商品批量上架接口入参
func (r *AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest) SetGoodsOnSaleCommand(_goodsOnSaleCommand *GoodsOnSaleCommandDto) error {
	r._goodsOnSaleCommand = _goodsOnSaleCommand
	r.Set("goods_on_sale_command", _goodsOnSaleCommand)
	return nil
}

// GetGoodsOnSaleCommand GoodsOnSaleCommand Getter
func (r AlibabaJymItemExternalGoodsBatchOnsaleAPIRequest) GetGoodsOnSaleCommand() *GoodsOnSaleCommandDto {
	return r._goodsOnSaleCommand
}
