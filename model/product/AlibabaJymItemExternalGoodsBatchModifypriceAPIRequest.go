package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymitemexternalgoodsbatchmodifypriceAPIRequest 交易猫外部商家批量商品改价接口 API请求
// alibaba.jym.item.external.goods.batch.modifyprice
//
// 供外部B端商家接入，提交批量商品改价请求，返回批量改价任务结果
type AlibabajymitemexternalgoodsbatchmodifypriceAPIRequest struct {
	model.Params
	// 商品批量改价接口入参
	_goodsPriceModifyCommand *GoodsPriceModifyCommandDto
}

// NewAlibabajymitemexternalgoodsbatchmodifypriceRequest 初始化AlibabajymitemexternalgoodsbatchmodifypriceAPIRequest对象
func NewAlibabajymitemexternalgoodsbatchmodifypriceRequest() *AlibabajymitemexternalgoodsbatchmodifypriceAPIRequest {
	return &AlibabajymitemexternalgoodsbatchmodifypriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymitemexternalgoodsbatchmodifypriceAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batch.modifyprice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymitemexternalgoodsbatchmodifypriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymitemexternalgoodsbatchmodifypriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsPriceModifyCommand is GoodsPriceModifyCommand Setter
// 商品批量改价接口入参
func (r *AlibabajymitemexternalgoodsbatchmodifypriceAPIRequest) SetGoodsPriceModifyCommand(_goodsPriceModifyCommand *GoodsPriceModifyCommandDto) error {
	r._goodsPriceModifyCommand = _goodsPriceModifyCommand
	r.Set("goods_price_modify_command", _goodsPriceModifyCommand)
	return nil
}

// GetGoodsPriceModifyCommand GoodsPriceModifyCommand Getter
func (r AlibabajymitemexternalgoodsbatchmodifypriceAPIRequest) GetGoodsPriceModifyCommand() *GoodsPriceModifyCommandDto {
	return r._goodsPriceModifyCommand
}
