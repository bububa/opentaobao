package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingDiagnosistitleAPIRequest 标题诊断 API请求
// alibaba.seaking.diagnosistitle
//
// 标题诊断
type AlibabaSeakingDiagnosistitleAPIRequest struct {
	model.Params
	// 类目id,没有的时候传-1
	_categoryId int64
	// 扩展信息
	_extra *Extra
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
}

// NewAlibabaSeakingDiagnosistitleRequest 初始化AlibabaSeakingDiagnosistitleAPIRequest对象
func NewAlibabaSeakingDiagnosistitleRequest() *AlibabaSeakingDiagnosistitleAPIRequest {
	return &AlibabaSeakingDiagnosistitleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSeakingDiagnosistitleAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.diagnosistitle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSeakingDiagnosistitleAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CategoryId Setter
// 类目id,没有的时候传-1
func (r *AlibabaSeakingDiagnosistitleAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// Get CategoryId Getter
func (r AlibabaSeakingDiagnosistitleAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// Set is Extra Setter
// 扩展信息
func (r *AlibabaSeakingDiagnosistitleAPIRequest) SetExtra(_extra *Extra) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// Get Extra Getter
func (r AlibabaSeakingDiagnosistitleAPIRequest) GetExtra() *Extra {
	return r._extra
}

// Set is Identifier Setter
// erp用户id
func (r *AlibabaSeakingDiagnosistitleAPIRequest) SetIdentifier(_identifier string) error {
	r._identifier = _identifier
	r.Set("identifier", _identifier)
	return nil
}

// Get Identifier Getter
func (r AlibabaSeakingDiagnosistitleAPIRequest) GetIdentifier() string {
	return r._identifier
}

// Set is IdentifierType Setter
// 调用来源(erp名称)
func (r *AlibabaSeakingDiagnosistitleAPIRequest) SetIdentifierType(_identifierType string) error {
	r._identifierType = _identifierType
	r.Set("identifier_type", _identifierType)
	return nil
}

// Get IdentifierType Getter
func (r AlibabaSeakingDiagnosistitleAPIRequest) GetIdentifierType() string {
	return r._identifierType
}

// Set is Language Setter
// 语种
func (r *AlibabaSeakingDiagnosistitleAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// Get Language Getter
func (r AlibabaSeakingDiagnosistitleAPIRequest) GetLanguage() string {
	return r._language
}

// Set is Platform Setter
// 商品所在平台（ae/icbu）
func (r *AlibabaSeakingDiagnosistitleAPIRequest) SetPlatform(_platform string) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// Get Platform Getter
func (r AlibabaSeakingDiagnosistitleAPIRequest) GetPlatform() string {
	return r._platform
}

// Set is Title Setter
// 标题
func (r *AlibabaSeakingDiagnosistitleAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// Get Title Getter
func (r AlibabaSeakingDiagnosistitleAPIRequest) GetTitle() string {
	return r._title
}
