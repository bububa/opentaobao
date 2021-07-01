package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRentItemAddAPIRequest
租赁商品发布 API请求
alibaba.idle.rent.item.add

发布闲鱼租赁商品 */
type AlibabaIdleRentItemAddAPIRequest struct {
	model.Params
	// 商品信息
	_paramRentItemDTO *RentItemDto
}

// NewAlibabaIdleRentItemAddRequest 初始化AlibabaIdleRentItemAddAPIRequest对象
func NewAlibabaIdleRentItemAddRequest() *AlibabaIdleRentItemAddAPIRequest {
	return &AlibabaIdleRentItemAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentItemAddAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.item.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentItemAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamRentItemDTO Setter
// 商品信息
func (r *AlibabaIdleRentItemAddAPIRequest) SetParamRentItemDTO(_paramRentItemDTO *RentItemDto) error {
	r._paramRentItemDTO = _paramRentItemDTO
	r.Set("param_rent_item_d_t_o", _paramRentItemDTO)
	return nil
}

// Get ParamRentItemDTO Getter
func (r AlibabaIdleRentItemAddAPIRequest) GetParamRentItemDTO() *RentItemDto {
	return r._paramRentItemDTO
}
