package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerentitemaddAPIRequest 租赁商品发布 API请求
// alibaba.idle.rent.item.add
//
// 发布闲鱼租赁商品
type AlibabaidlerentitemaddAPIRequest struct {
	model.Params
	// 商品信息
	_paramRentItemDTO *RentItemDto
}

// NewAlibabaidlerentitemaddRequest 初始化AlibabaidlerentitemaddAPIRequest对象
func NewAlibabaidlerentitemaddRequest() *AlibabaidlerentitemaddAPIRequest {
	return &AlibabaidlerentitemaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerentitemaddAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.item.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerentitemaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerentitemaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRentItemDTO is ParamRentItemDTO Setter
// 商品信息
func (r *AlibabaidlerentitemaddAPIRequest) SetParamRentItemDTO(_paramRentItemDTO *RentItemDto) error {
	r._paramRentItemDTO = _paramRentItemDTO
	r.Set("param_rent_item_d_t_o", _paramRentItemDTO)
	return nil
}

// GetParamRentItemDTO ParamRentItemDTO Getter
func (r AlibabaidlerentitemaddAPIRequest) GetParamRentItemDTO() *RentItemDto {
	return r._paramRentItemDTO
}
