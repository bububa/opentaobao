package icbu

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuCategoryLevelAttrGetAPIRequest 层级属性的子属性获取 API请求
// alibaba.icbu.category.level.attr.get
//
// 用于获取层级属性（车型库）的子属性和属性值
type AlibabaIcbuCategoryLevelAttrGetAPIRequest struct {
	model.Params
	// 属性值request对象
	_attributeValueRequest *LevelAttributeValueRequest
}

// NewAlibabaIcbuCategoryLevelAttrGetRequest 初始化AlibabaIcbuCategoryLevelAttrGetAPIRequest对象
func NewAlibabaIcbuCategoryLevelAttrGetRequest() *AlibabaIcbuCategoryLevelAttrGetAPIRequest {
	return &AlibabaIcbuCategoryLevelAttrGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIcbuCategoryLevelAttrGetAPIRequest) Reset() {
	r._attributeValueRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategoryLevelAttrGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.level.attr.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategoryLevelAttrGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuCategoryLevelAttrGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAttributeValueRequest is AttributeValueRequest Setter
// 属性值request对象
func (r *AlibabaIcbuCategoryLevelAttrGetAPIRequest) SetAttributeValueRequest(_attributeValueRequest *LevelAttributeValueRequest) error {
	r._attributeValueRequest = _attributeValueRequest
	r.Set("attribute_value_request", _attributeValueRequest)
	return nil
}

// GetAttributeValueRequest AttributeValueRequest Getter
func (r AlibabaIcbuCategoryLevelAttrGetAPIRequest) GetAttributeValueRequest() *LevelAttributeValueRequest {
	return r._attributeValueRequest
}

var poolAlibabaIcbuCategoryLevelAttrGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIcbuCategoryLevelAttrGetRequest()
	},
}

// GetAlibabaIcbuCategoryLevelAttrGetRequest 从 sync.Pool 获取 AlibabaIcbuCategoryLevelAttrGetAPIRequest
func GetAlibabaIcbuCategoryLevelAttrGetAPIRequest() *AlibabaIcbuCategoryLevelAttrGetAPIRequest {
	return poolAlibabaIcbuCategoryLevelAttrGetAPIRequest.Get().(*AlibabaIcbuCategoryLevelAttrGetAPIRequest)
}

// ReleaseAlibabaIcbuCategoryLevelAttrGetAPIRequest 将 AlibabaIcbuCategoryLevelAttrGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaIcbuCategoryLevelAttrGetAPIRequest(v *AlibabaIcbuCategoryLevelAttrGetAPIRequest) {
	v.Reset()
	poolAlibabaIcbuCategoryLevelAttrGetAPIRequest.Put(v)
}
