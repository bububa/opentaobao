package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaseakingdiagnosistitleAPIRequest 标题诊断 API请求
// alibaba.seaking.diagnosistitle
//
// 标题诊断
type AlibabaseakingdiagnosistitleAPIRequest struct {
	model.Params
	// erp用户id
	_identifier string
	// 调用来源(erp名称)
	_identifierType string
	// 语种
	_language string
	// 商品所在平台（ae/icbu）
	_platform string
	// 标题
	_title string
	// 类目id,没有的时候传-1
	_categoryId int64
	// 扩展信息
	_extra *Extra
}

// NewAlibabaseakingdiagnosistitleRequest 初始化AlibabaseakingdiagnosistitleAPIRequest对象
func NewAlibabaseakingdiagnosistitleRequest() *AlibabaseakingdiagnosistitleAPIRequest {
	return &AlibabaseakingdiagnosistitleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaseakingdiagnosistitleAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.diagnosistitle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaseakingdiagnosistitleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaseakingdiagnosistitleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIdentifier is Identifier Setter
// erp用户id
func (r *AlibabaseakingdiagnosistitleAPIRequest) SetIdentifier(_identifier string) error {
	r._identifier = _identifier
	r.Set("identifier", _identifier)
	return nil
}

// GetIdentifier Identifier Getter
func (r AlibabaseakingdiagnosistitleAPIRequest) GetIdentifier() string {
	return r._identifier
}

// SetIdentifierType is IdentifierType Setter
// 调用来源(erp名称)
func (r *AlibabaseakingdiagnosistitleAPIRequest) SetIdentifierType(_identifierType string) error {
	r._identifierType = _identifierType
	r.Set("identifier_type", _identifierType)
	return nil
}

// GetIdentifierType IdentifierType Getter
func (r AlibabaseakingdiagnosistitleAPIRequest) GetIdentifierType() string {
	return r._identifierType
}

// SetLanguage is Language Setter
// 语种
func (r *AlibabaseakingdiagnosistitleAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AlibabaseakingdiagnosistitleAPIRequest) GetLanguage() string {
	return r._language
}

// SetPlatform is Platform Setter
// 商品所在平台（ae/icbu）
func (r *AlibabaseakingdiagnosistitleAPIRequest) SetPlatform(_platform string) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r AlibabaseakingdiagnosistitleAPIRequest) GetPlatform() string {
	return r._platform
}

// SetTitle is Title Setter
// 标题
func (r *AlibabaseakingdiagnosistitleAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r AlibabaseakingdiagnosistitleAPIRequest) GetTitle() string {
	return r._title
}

// SetCategoryId is CategoryId Setter
// 类目id,没有的时候传-1
func (r *AlibabaseakingdiagnosistitleAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r AlibabaseakingdiagnosistitleAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetExtra is Extra Setter
// 扩展信息
func (r *AlibabaseakingdiagnosistitleAPIRequest) SetExtra(_extra *Extra) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r AlibabaseakingdiagnosistitleAPIRequest) GetExtra() *Extra {
	return r._extra
}
