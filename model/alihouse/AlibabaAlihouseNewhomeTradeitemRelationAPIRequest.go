package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeTradeitemRelationAPIRequest 货独立绑定货品 API请求
// alibaba.alihouse.newhome.tradeitem.relation
//
// 货独立绑定货品
type AlibabaAlihouseNewhomeTradeitemRelationAPIRequest struct {
	model.Params
	// 货独立绑定货品关系请求体
	_relationBindingDto *RelationBindingDto
}

// NewAlibabaAlihouseNewhomeTradeitemRelationRequest 初始化AlibabaAlihouseNewhomeTradeitemRelationAPIRequest对象
func NewAlibabaAlihouseNewhomeTradeitemRelationRequest() *AlibabaAlihouseNewhomeTradeitemRelationAPIRequest {
	return &AlibabaAlihouseNewhomeTradeitemRelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeTradeitemRelationAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.tradeitem.relation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeTradeitemRelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeTradeitemRelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRelationBindingDto is RelationBindingDto Setter
// 货独立绑定货品关系请求体
func (r *AlibabaAlihouseNewhomeTradeitemRelationAPIRequest) SetRelationBindingDto(_relationBindingDto *RelationBindingDto) error {
	r._relationBindingDto = _relationBindingDto
	r.Set("relation_binding_dto", _relationBindingDto)
	return nil
}

// GetRelationBindingDto RelationBindingDto Getter
func (r AlibabaAlihouseNewhomeTradeitemRelationAPIRequest) GetRelationBindingDto() *RelationBindingDto {
	return r._relationBindingDto
}
