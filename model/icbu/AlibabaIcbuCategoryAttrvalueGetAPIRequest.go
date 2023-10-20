package icbu

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuCategoryAttrvalueGetAPIRequest) Reset() {
	r._attributeValueRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategoryAttrvalueGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.attrvalue.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategoryAttrvalueGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuCategoryAttrvalueGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaIcbuCategoryAttrvalueGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuCategoryAttrvalueGetRequest()
	},
}

// GetAlibabaIcbuCategoryAttrvalueGetRequest 从 sync.Pool 获取 AlibabaIcbuCategoryAttrvalueGetAPIRequest
func GetAlibabaIcbuCategoryAttrvalueGetAPIRequest() *AlibabaIcbuCategoryAttrvalueGetAPIRequest {
	return poolAlibabaIcbuCategoryAttrvalueGetAPIRequest.Get().(*AlibabaIcbuCategoryAttrvalueGetAPIRequest)
}

// ReleaseAlibabaIcbuCategoryAttrvalueGetAPIRequest 将 AlibabaIcbuCategoryAttrvalueGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuCategoryAttrvalueGetAPIRequest(v *AlibabaIcbuCategoryAttrvalueGetAPIRequest) {
	v.Reset()
	poolAlibabaIcbuCategoryAttrvalueGetAPIRequest.Put(v)
}
