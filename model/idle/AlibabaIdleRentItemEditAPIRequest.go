package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerentitemeditAPIRequest 租赁商品编辑 API请求
// alibaba.idle.rent.item.edit
//
// 发布闲鱼租赁商品
type AlibabaidlerentitemeditAPIRequest struct {
	model.Params
	// 商品信息
	_paramRentItemDTO *RentItemDto
}

// NewAlibabaidlerentitemeditRequest 初始化AlibabaidlerentitemeditAPIRequest对象
func NewAlibabaidlerentitemeditRequest() *AlibabaidlerentitemeditAPIRequest {
	return &AlibabaidlerentitemeditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerentitemeditAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.item.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerentitemeditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerentitemeditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamRentItemDTO is ParamRentItemDTO Setter
// 商品信息
func (r *AlibabaidlerentitemeditAPIRequest) SetParamRentItemDTO(_paramRentItemDTO *RentItemDto) error {
	r._paramRentItemDTO = _paramRentItemDTO
	r.Set("param_rent_item_d_t_o", _paramRentItemDTO)
	return nil
}

// GetParamRentItemDTO ParamRentItemDTO Getter
func (r AlibabaidlerentitemeditAPIRequest) GetParamRentItemDTO() *RentItemDto {
	return r._paramRentItemDTO
}
