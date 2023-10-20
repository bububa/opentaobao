package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbucategoryattrvaluegetAPIRequest 属性值获取 API请求
// alibaba.icbu.category.attrvalue.get
//
// 属性值获取
type AlibabaicbucategoryattrvaluegetAPIRequest struct {
	model.Params
	// 属性值request对象
	_attributeValueRequest *AttributeValueRequest
}

// NewAlibabaicbucategoryattrvaluegetRequest 初始化AlibabaicbucategoryattrvaluegetAPIRequest对象
func NewAlibabaicbucategoryattrvaluegetRequest() *AlibabaicbucategoryattrvaluegetAPIRequest {
	return &AlibabaicbucategoryattrvaluegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbucategoryattrvaluegetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.attrvalue.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbucategoryattrvaluegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbucategoryattrvaluegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAttributeValueRequest is AttributeValueRequest Setter
// 属性值request对象
func (r *AlibabaicbucategoryattrvaluegetAPIRequest) SetAttributeValueRequest(_attributeValueRequest *AttributeValueRequest) error {
	r._attributeValueRequest = _attributeValueRequest
	r.Set("attribute_value_request", _attributeValueRequest)
	return nil
}

// GetAttributeValueRequest AttributeValueRequest Getter
func (r AlibabaicbucategoryattrvaluegetAPIRequest) GetAttributeValueRequest() *AttributeValueRequest {
	return r._attributeValueRequest
}
