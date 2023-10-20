package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymitemexternalgoodsbatchonsaleAPIRequest 交易猫外部商家批量上架商品接口 API请求
// alibaba.jym.item.external.goods.batch.onsale
//
// 供外部B端商家接入，提交批量上架商品请求，返回批量上架任务结果
type AlibabajymitemexternalgoodsbatchonsaleAPIRequest struct {
	model.Params
	// 商品批量上架接口入参
	_goodsOnSaleCommand *GoodsOnSaleCommandDto
}

// NewAlibabajymitemexternalgoodsbatchonsaleRequest 初始化AlibabajymitemexternalgoodsbatchonsaleAPIRequest对象
func NewAlibabajymitemexternalgoodsbatchonsaleRequest() *AlibabajymitemexternalgoodsbatchonsaleAPIRequest {
	return &AlibabajymitemexternalgoodsbatchonsaleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymitemexternalgoodsbatchonsaleAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batch.onsale"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymitemexternalgoodsbatchonsaleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymitemexternalgoodsbatchonsaleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsOnSaleCommand is GoodsOnSaleCommand Setter
// 商品批量上架接口入参
func (r *AlibabajymitemexternalgoodsbatchonsaleAPIRequest) SetGoodsOnSaleCommand(_goodsOnSaleCommand *GoodsOnSaleCommandDto) error {
	r._goodsOnSaleCommand = _goodsOnSaleCommand
	r.Set("goods_on_sale_command", _goodsOnSaleCommand)
	return nil
}

// GetGoodsOnSaleCommand GoodsOnSaleCommand Getter
func (r AlibabajymitemexternalgoodsbatchonsaleAPIRequest) GetGoodsOnSaleCommand() *GoodsOnSaleCommandDto {
	return r._goodsOnSaleCommand
}
