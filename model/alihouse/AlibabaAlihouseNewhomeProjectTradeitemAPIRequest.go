package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojecttradeitemAPIRequest 新二租品同步 API请求
// alibaba.alihouse.newhome.project.tradeitem
//
// 新二品同步
type AlibabaalihousenewhomeprojecttradeitemAPIRequest struct {
	model.Params
	// 请求对象
	_goodsDto *GoodsDto
}

// NewAlibabaalihousenewhomeprojecttradeitemRequest 初始化AlibabaalihousenewhomeprojecttradeitemAPIRequest对象
func NewAlibabaalihousenewhomeprojecttradeitemRequest() *AlibabaalihousenewhomeprojecttradeitemAPIRequest {
	return &AlibabaalihousenewhomeprojecttradeitemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojecttradeitemAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.tradeitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojecttradeitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojecttradeitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGoodsDto is GoodsDto Setter
// 请求对象
func (r *AlibabaalihousenewhomeprojecttradeitemAPIRequest) SetGoodsDto(_goodsDto *GoodsDto) error {
	r._goodsDto = _goodsDto
	r.Set("goods_dto", _goodsDto)
	return nil
}

// GetGoodsDto GoodsDto Getter
func (r AlibabaalihousenewhomeprojecttradeitemAPIRequest) GetGoodsDto() *GoodsDto {
	return r._goodsDto
}
