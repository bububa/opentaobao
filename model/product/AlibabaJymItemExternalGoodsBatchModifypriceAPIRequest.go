package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest 交易猫外部商家批量商品改价接口 API请求
// alibaba.jym.item.external.goods.batch.modifyprice
//
// 供外部B端商家接入，提交批量商品改价请求，返回批量改价任务结果
type AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest struct {
	model.Params
	// 商品批量改价接口入参
	_goodsPriceModifyCommand *GoodsPriceModifyCommandDto
}

// NewAlibabaJymItemExternalGoodsBatchModifypriceRequest 初始化AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest对象
func NewAlibabaJymItemExternalGoodsBatchModifypriceRequest() *AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest {
	return &AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batch.modifyprice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetGoodsPriceModifyCommand is GoodsPriceModifyCommand Setter
// 商品批量改价接口入参
func (r *AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest) SetGoodsPriceModifyCommand(_goodsPriceModifyCommand *GoodsPriceModifyCommandDto) error {
	r._goodsPriceModifyCommand = _goodsPriceModifyCommand
	r.Set("goods_price_modify_command", _goodsPriceModifyCommand)
	return nil
}

// GetGoodsPriceModifyCommand GoodsPriceModifyCommand Getter
func (r AlibabaJymItemExternalGoodsBatchModifypriceAPIRequest) GetGoodsPriceModifyCommand() *GoodsPriceModifyCommandDto {
	return r._goodsPriceModifyCommand
}
