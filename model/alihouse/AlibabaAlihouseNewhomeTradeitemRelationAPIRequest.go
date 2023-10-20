package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhometradeitemrelationAPIRequest 货独立绑定货品 API请求
// alibaba.alihouse.newhome.tradeitem.relation
//
// 货独立绑定货品
type AlibabaalihousenewhometradeitemrelationAPIRequest struct {
	model.Params
	// 货独立绑定货品关系请求体
	_relationBindingDto *RelationBindingDto
}

// NewAlibabaalihousenewhometradeitemrelationRequest 初始化AlibabaalihousenewhometradeitemrelationAPIRequest对象
func NewAlibabaalihousenewhometradeitemrelationRequest() *AlibabaalihousenewhometradeitemrelationAPIRequest {
	return &AlibabaalihousenewhometradeitemrelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhometradeitemrelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.tradeitem.relation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhometradeitemrelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhometradeitemrelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationBindingDto is RelationBindingDto Setter
// 货独立绑定货品关系请求体
func (r *AlibabaalihousenewhometradeitemrelationAPIRequest) SetRelationBindingDto(_relationBindingDto *RelationBindingDto) error {
	r._relationBindingDto = _relationBindingDto
	r.Set("relation_binding_dto", _relationBindingDto)
	return nil
}

// GetRelationBindingDto RelationBindingDto Getter
func (r AlibabaalihousenewhometradeitemrelationAPIRequest) GetRelationBindingDto() *RelationBindingDto {
	return r._relationBindingDto
}
