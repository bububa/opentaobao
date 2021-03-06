package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingAititlegenerateAPIRequest 标题智能优化 API请求
// alibaba.seaking.aititlegenerate
//
// 标题智能优化
type AlibabaSeakingAititlegenerateAPIRequest struct {
	model.Params
	// erp用户id
	_identifier string
	// 语种
	_language string
	// 商品属性
	_attributes string
	// 调用来源(erp名称)
	_identifierType string
	// 标题
	_title string
	// 商品所在平台（ae/icbu）
	_platform string
	// 扩展信息
	_extra *Extra
	// 类目id,没有的时候传-1
	_categoryId int64
}

// NewAlibabaSeakingAititlegenerateRequest 初始化AlibabaSeakingAititlegenerateAPIRequest对象
func NewAlibabaSeakingAititlegenerateRequest() *AlibabaSeakingAititlegenerateAPIRequest {
	return &AlibabaSeakingAititlegenerateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSeakingAititlegenerateAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.aititlegenerate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSeakingAititlegenerateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetIdentifier is Identifier Setter
// erp用户id
func (r *AlibabaSeakingAititlegenerateAPIRequest) SetIdentifier(_identifier string) error {
	r._identifier = _identifier
	r.Set("identifier", _identifier)
	return nil
}

// GetIdentifier Identifier Getter
func (r AlibabaSeakingAititlegenerateAPIRequest) GetIdentifier() string {
	return r._identifier
}

// SetLanguage is Language Setter
// 语种
func (r *AlibabaSeakingAititlegenerateAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AlibabaSeakingAititlegenerateAPIRequest) GetLanguage() string {
	return r._language
}

// SetAttributes is Attributes Setter
// 商品属性
func (r *AlibabaSeakingAititlegenerateAPIRequest) SetAttributes(_attributes string) error {
	r._attributes = _attributes
	r.Set("attributes", _attributes)
	return nil
}

// GetAttributes Attributes Getter
func (r AlibabaSeakingAititlegenerateAPIRequest) GetAttributes() string {
	return r._attributes
}

// SetIdentifierType is IdentifierType Setter
// 调用来源(erp名称)
func (r *AlibabaSeakingAititlegenerateAPIRequest) SetIdentifierType(_identifierType string) error {
	r._identifierType = _identifierType
	r.Set("identifier_type", _identifierType)
	return nil
}

// GetIdentifierType IdentifierType Getter
func (r AlibabaSeakingAititlegenerateAPIRequest) GetIdentifierType() string {
	return r._identifierType
}

// SetTitle is Title Setter
// 标题
func (r *AlibabaSeakingAititlegenerateAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r AlibabaSeakingAititlegenerateAPIRequest) GetTitle() string {
	return r._title
}

// SetPlatform is Platform Setter
// 商品所在平台（ae/icbu）
func (r *AlibabaSeakingAititlegenerateAPIRequest) SetPlatform(_platform string) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r AlibabaSeakingAititlegenerateAPIRequest) GetPlatform() string {
	return r._platform
}

// SetExtra is Extra Setter
// 扩展信息
func (r *AlibabaSeakingAititlegenerateAPIRequest) SetExtra(_extra *Extra) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r AlibabaSeakingAititlegenerateAPIRequest) GetExtra() *Extra {
	return r._extra
}

// SetCategoryId is CategoryId Setter
// 类目id,没有的时候传-1
func (r *AlibabaSeakingAititlegenerateAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r AlibabaSeakingAititlegenerateAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}
