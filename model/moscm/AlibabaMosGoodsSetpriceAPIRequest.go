package moscm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosgoodssetpriceAPIRequest 价格变更接口 API请求
// alibaba.mos.goods.setprice
//
// 价格变更接口，供供应商修改价格时使用。
type AlibabamosgoodssetpriceAPIRequest struct {
	model.Params
	// 价格变更对象列表
	_priceDtoList []PriceDto
}

// NewAlibabamosgoodssetpriceRequest 初始化AlibabamosgoodssetpriceAPIRequest对象
func NewAlibabamosgoodssetpriceRequest() *AlibabamosgoodssetpriceAPIRequest {
	return &AlibabamosgoodssetpriceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamosgoodssetpriceAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.goods.setprice"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamosgoodssetpriceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamosgoodssetpriceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPriceDtoList is PriceDtoList Setter
// 价格变更对象列表
func (r *AlibabamosgoodssetpriceAPIRequest) SetPriceDtoList(_priceDtoList []PriceDto) error {
	r._priceDtoList = _priceDtoList
	r.Set("price_dto_list", _priceDtoList)
	return nil
}

// GetPriceDtoList PriceDtoList Getter
func (r AlibabamosgoodssetpriceAPIRequest) GetPriceDtoList() []PriceDto {
	return r._priceDtoList
}
