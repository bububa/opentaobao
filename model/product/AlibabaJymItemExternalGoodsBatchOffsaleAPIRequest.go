package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymitemexternalgoodsbatchoffsaleAPIRequest 交易猫外部商家批量下架商品接口 API请求
// alibaba.jym.item.external.goods.batch.offsale
//
// 供外部B端商家接入，提交批量下架商品请求，返回批量下架任务结果
type AlibabajymitemexternalgoodsbatchoffsaleAPIRequest struct {
	model.Params
	// 商品批量下架接口入参
	_goodsOffSaleCommand *GoodsOffSaleCommandDto
}

// NewAlibabajymitemexternalgoodsbatchoffsaleRequest 初始化AlibabajymitemexternalgoodsbatchoffsaleAPIRequest对象
func NewAlibabajymitemexternalgoodsbatchoffsaleRequest() *AlibabajymitemexternalgoodsbatchoffsaleAPIRequest {
	return &AlibabajymitemexternalgoodsbatchoffsaleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymitemexternalgoodsbatchoffsaleAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batch.offsale"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymitemexternalgoodsbatchoffsaleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymitemexternalgoodsbatchoffsaleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsOffSaleCommand is GoodsOffSaleCommand Setter
// 商品批量下架接口入参
func (r *AlibabajymitemexternalgoodsbatchoffsaleAPIRequest) SetGoodsOffSaleCommand(_goodsOffSaleCommand *GoodsOffSaleCommandDto) error {
	r._goodsOffSaleCommand = _goodsOffSaleCommand
	r.Set("goods_off_sale_command", _goodsOffSaleCommand)
	return nil
}

// GetGoodsOffSaleCommand GoodsOffSaleCommand Getter
func (r AlibabajymitemexternalgoodsbatchoffsaleAPIRequest) GetGoodsOffSaleCommand() *GoodsOffSaleCommandDto {
	return r._goodsOffSaleCommand
}
