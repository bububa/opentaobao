package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest 交易猫外部商家批量下架商品接口 API请求
// alibaba.jym.item.external.goods.batch.offsale
//
// 供外部B端商家接入，提交批量下架商品请求，返回批量下架任务结果
type AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest struct {
	model.Params
	// 商品批量下架接口入参
	_goodsOffSaleCommand *GoodsOffSaleCommandDto
}

// NewAlibabaJymItemExternalGoodsBatchOffsaleRequest 初始化AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest对象
func NewAlibabaJymItemExternalGoodsBatchOffsaleRequest() *AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest {
	return &AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batch.offsale"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetGoodsOffSaleCommand is GoodsOffSaleCommand Setter
// 商品批量下架接口入参
func (r *AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest) SetGoodsOffSaleCommand(_goodsOffSaleCommand *GoodsOffSaleCommandDto) error {
	r._goodsOffSaleCommand = _goodsOffSaleCommand
	r.Set("goods_off_sale_command", _goodsOffSaleCommand)
	return nil
}

// GetGoodsOffSaleCommand GoodsOffSaleCommand Getter
func (r AlibabaJymItemExternalGoodsBatchOffsaleAPIRequest) GetGoodsOffSaleCommand() *GoodsOffSaleCommandDto {
	return r._goodsOffSaleCommand
}
