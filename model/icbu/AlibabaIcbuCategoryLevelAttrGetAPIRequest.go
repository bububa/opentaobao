package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbucategorylevelattrgetAPIRequest 层级属性的子属性获取 API请求
// alibaba.icbu.category.level.attr.get
//
// 用于获取层级属性（车型库）的子属性和属性值
type AlibabaicbucategorylevelattrgetAPIRequest struct {
	model.Params
	// 属性值request对象
	_attributeValueRequest *LevelAttributeValueRequest
}

// NewAlibabaicbucategorylevelattrgetRequest 初始化AlibabaicbucategorylevelattrgetAPIRequest对象
func NewAlibabaicbucategorylevelattrgetRequest() *AlibabaicbucategorylevelattrgetAPIRequest {
	return &AlibabaicbucategorylevelattrgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbucategorylevelattrgetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.level.attr.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbucategorylevelattrgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbucategorylevelattrgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAttributeValueRequest is AttributeValueRequest Setter
// 属性值request对象
func (r *AlibabaicbucategorylevelattrgetAPIRequest) SetAttributeValueRequest(_attributeValueRequest *LevelAttributeValueRequest) error {
	r._attributeValueRequest = _attributeValueRequest
	r.Set("attribute_value_request", _attributeValueRequest)
	return nil
}

// GetAttributeValueRequest AttributeValueRequest Getter
func (r AlibabaicbucategorylevelattrgetAPIRequest) GetAttributeValueRequest() *LevelAttributeValueRequest {
	return r._attributeValueRequest
}
