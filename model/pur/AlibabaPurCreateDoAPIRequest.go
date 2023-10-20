package pur

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapurcreatedoAPIRequest top创建DO/RT接口 API请求
// alibaba.pur.create.do
//
// 创建发货单,先创建DO，异步创建RT
type AlibabapurcreatedoAPIRequest struct {
	model.Params
	// 发货单
	_deliveryTopDTO *DeliveryTopDto
}

// NewAlibabapurcreatedoRequest 初始化AlibabapurcreatedoAPIRequest对象
func NewAlibabapurcreatedoRequest() *AlibabapurcreatedoAPIRequest {
	return &AlibabapurcreatedoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapurcreatedoAPIRequest) GetApiMethodName() string {
	return "alibaba.pur.create.do"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapurcreatedoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapurcreatedoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryTopDTO is DeliveryTopDTO Setter
// 发货单
func (r *AlibabapurcreatedoAPIRequest) SetDeliveryTopDTO(_deliveryTopDTO *DeliveryTopDto) error {
	r._deliveryTopDTO = _deliveryTopDTO
	r.Set("delivery_top_d_t_o", _deliveryTopDTO)
	return nil
}

// GetDeliveryTopDTO DeliveryTopDTO Getter
func (r AlibabapurcreatedoAPIRequest) GetDeliveryTopDTO() *DeliveryTopDto {
	return r._deliveryTopDTO
}
