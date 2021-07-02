package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingImagetranslateAPIRequest 图片机器翻译 API请求
// alibaba.seaking.imagetranslate
//
// 图片机器翻译
type AlibabaSeakingImagetranslateAPIRequest struct {
	model.Params
	// erp用户id
	_identifier string
	// 目标语种
	_targetLang string
	// 源语种
	_sourceLang string
	// 调用来源(erp名称)
	_identifierType string
	// 原图url
	_url string
	// 扩展信息
	_extra *Extra
}

// NewAlibabaSeakingImagetranslateRequest 初始化AlibabaSeakingImagetranslateAPIRequest对象
func NewAlibabaSeakingImagetranslateRequest() *AlibabaSeakingImagetranslateAPIRequest {
	return &AlibabaSeakingImagetranslateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSeakingImagetranslateAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.imagetranslate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSeakingImagetranslateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetIdentifier is Identifier Setter
// erp用户id
func (r *AlibabaSeakingImagetranslateAPIRequest) SetIdentifier(_identifier string) error {
	r._identifier = _identifier
	r.Set("identifier", _identifier)
	return nil
}

// GetIdentifier Identifier Getter
func (r AlibabaSeakingImagetranslateAPIRequest) GetIdentifier() string {
	return r._identifier
}

// SetTargetLang is TargetLang Setter
// 目标语种
func (r *AlibabaSeakingImagetranslateAPIRequest) SetTargetLang(_targetLang string) error {
	r._targetLang = _targetLang
	r.Set("target_lang", _targetLang)
	return nil
}

// GetTargetLang TargetLang Getter
func (r AlibabaSeakingImagetranslateAPIRequest) GetTargetLang() string {
	return r._targetLang
}

// SetSourceLang is SourceLang Setter
// 源语种
func (r *AlibabaSeakingImagetranslateAPIRequest) SetSourceLang(_sourceLang string) error {
	r._sourceLang = _sourceLang
	r.Set("source_lang", _sourceLang)
	return nil
}

// GetSourceLang SourceLang Getter
func (r AlibabaSeakingImagetranslateAPIRequest) GetSourceLang() string {
	return r._sourceLang
}

// SetIdentifierType is IdentifierType Setter
// 调用来源(erp名称)
func (r *AlibabaSeakingImagetranslateAPIRequest) SetIdentifierType(_identifierType string) error {
	r._identifierType = _identifierType
	r.Set("identifier_type", _identifierType)
	return nil
}

// GetIdentifierType IdentifierType Getter
func (r AlibabaSeakingImagetranslateAPIRequest) GetIdentifierType() string {
	return r._identifierType
}

// SetUrl is Url Setter
// 原图url
func (r *AlibabaSeakingImagetranslateAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r AlibabaSeakingImagetranslateAPIRequest) GetUrl() string {
	return r._url
}

// SetExtra is Extra Setter
// 扩展信息
func (r *AlibabaSeakingImagetranslateAPIRequest) SetExtra(_extra *Extra) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r AlibabaSeakingImagetranslateAPIRequest) GetExtra() *Extra {
	return r._extra
}
