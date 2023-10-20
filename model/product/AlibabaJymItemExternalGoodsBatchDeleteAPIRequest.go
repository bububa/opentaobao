package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymitemexternalgoodsbatchdeleteAPIRequest 交易猫外部商家批量删除商品接口 API请求
// alibaba.jym.item.external.goods.batch.delete
//
// 交易猫外部商家批量删除商品接口
type AlibabajymitemexternalgoodsbatchdeleteAPIRequest struct {
	model.Params
	// 商品批量删除接口入参
	_goodsDeleteCommandDto *GoodsDeleteCommandDto
}

// NewAlibabajymitemexternalgoodsbatchdeleteRequest 初始化AlibabajymitemexternalgoodsbatchdeleteAPIRequest对象
func NewAlibabajymitemexternalgoodsbatchdeleteRequest() *AlibabajymitemexternalgoodsbatchdeleteAPIRequest {
	return &AlibabajymitemexternalgoodsbatchdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymitemexternalgoodsbatchdeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.external.goods.batch.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymitemexternalgoodsbatchdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymitemexternalgoodsbatchdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsDeleteCommandDto is GoodsDeleteCommandDto Setter
// 商品批量删除接口入参
func (r *AlibabajymitemexternalgoodsbatchdeleteAPIRequest) SetGoodsDeleteCommandDto(_goodsDeleteCommandDto *GoodsDeleteCommandDto) error {
	r._goodsDeleteCommandDto = _goodsDeleteCommandDto
	r.Set("goods_delete_command_dto", _goodsDeleteCommandDto)
	return nil
}

// GetGoodsDeleteCommandDto GoodsDeleteCommandDto Getter
func (r AlibabajymitemexternalgoodsbatchdeleteAPIRequest) GetGoodsDeleteCommandDto() *GoodsDeleteCommandDto {
	return r._goodsDeleteCommandDto
}
