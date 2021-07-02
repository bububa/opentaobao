package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuCategoryAttrvalueGetAPIRequest 属性值获取 API请求
// alibaba.icbu.category.attrvalue.get
//
// 属性值获取
type AlibabaIcbuCategoryAttrvalueGetAPIRequest struct {
	model.Params
	// 属性值request对象
	_attributeValueRequest *AttributeValueRequest
}

// NewAlibabaIcbuCategoryAttrvalueGetRequest 初始化AlibabaIcbuCategoryAttrvalueGetAPIRequest对象
func NewAlibabaIcbuCategoryAttrvalueGetRequest() *AlibabaIcbuCategoryAttrvalueGetAPIRequest {
	return &AlibabaIcbuCategoryAttrvalueGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategoryAttrvalueGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.attrvalue.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategoryAttrvalueGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAttributeValueRequest is AttributeValueRequest Setter
// 属性值request对象
func (r *AlibabaIcbuCategoryAttrvalueGetAPIRequest) SetAttributeValueRequest(_attributeValueRequest *AttributeValueRequest) error {
	r._attributeValueRequest = _attributeValueRequest
	r.Set("attribute_value_request", _attributeValueRequest)
	return nil
}

// GetAttributeValueRequest AttributeValueRequest Getter
func (r AlibabaIcbuCategoryAttrvalueGetAPIRequest) GetAttributeValueRequest() *AttributeValueRequest {
	return r._attributeValueRequest
}
