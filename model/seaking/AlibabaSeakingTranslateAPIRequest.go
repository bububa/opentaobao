package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
MT定制接口 API请求
alibaba.seaking.translate

MT定制接口
*/
type AlibabaSeakingTranslateAPIRequest struct {
    model.Params
    // 定制用户id
    _identifier   string
    // 目标语种
    _targetLang   string
    // 源语种
    _sourceLang   string
    // 原文
    _sourceText   string
    // 原文格式(text/html)
    _sourceFormat   string
    // 定制用户类型
    _identifierType   string
    // 原文类型(title: 标题/offer: 详描/message: 消息)
    _fieldType   string
    // 扩展信息
    _extra   *Extra
}

// 初始化AlibabaSeakingTranslateAPIRequest对象
func NewAlibabaSeakingTranslateRequest() *AlibabaSeakingTranslateAPIRequest{
    return &AlibabaSeakingTranslateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingTranslateAPIRequest) GetApiMethodName() string {
    return "alibaba.seaking.translate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingTranslateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Identifier Setter
// 定制用户id
func (r *AlibabaSeakingTranslateAPIRequest) SetIdentifier(_identifier string) error {
    r._identifier = _identifier
    r.Set("identifier", _identifier)
    return nil
}

// Identifier Getter
func (r AlibabaSeakingTranslateAPIRequest) GetIdentifier() string {
    return r._identifier
}
// TargetLang Setter
// 目标语种
func (r *AlibabaSeakingTranslateAPIRequest) SetTargetLang(_targetLang string) error {
    r._targetLang = _targetLang
    r.Set("target_lang", _targetLang)
    return nil
}

// TargetLang Getter
func (r AlibabaSeakingTranslateAPIRequest) GetTargetLang() string {
    return r._targetLang
}
// SourceLang Setter
// 源语种
func (r *AlibabaSeakingTranslateAPIRequest) SetSourceLang(_sourceLang string) error {
    r._sourceLang = _sourceLang
    r.Set("source_lang", _sourceLang)
    return nil
}

// SourceLang Getter
func (r AlibabaSeakingTranslateAPIRequest) GetSourceLang() string {
    return r._sourceLang
}
// SourceText Setter
// 原文
func (r *AlibabaSeakingTranslateAPIRequest) SetSourceText(_sourceText string) error {
    r._sourceText = _sourceText
    r.Set("source_text", _sourceText)
    return nil
}

// SourceText Getter
func (r AlibabaSeakingTranslateAPIRequest) GetSourceText() string {
    return r._sourceText
}
// SourceFormat Setter
// 原文格式(text/html)
func (r *AlibabaSeakingTranslateAPIRequest) SetSourceFormat(_sourceFormat string) error {
    r._sourceFormat = _sourceFormat
    r.Set("source_format", _sourceFormat)
    return nil
}

// SourceFormat Getter
func (r AlibabaSeakingTranslateAPIRequest) GetSourceFormat() string {
    return r._sourceFormat
}
// IdentifierType Setter
// 定制用户类型
func (r *AlibabaSeakingTranslateAPIRequest) SetIdentifierType(_identifierType string) error {
    r._identifierType = _identifierType
    r.Set("identifier_type", _identifierType)
    return nil
}

// IdentifierType Getter
func (r AlibabaSeakingTranslateAPIRequest) GetIdentifierType() string {
    return r._identifierType
}
// FieldType Setter
// 原文类型(title: 标题/offer: 详描/message: 消息)
func (r *AlibabaSeakingTranslateAPIRequest) SetFieldType(_fieldType string) error {
    r._fieldType = _fieldType
    r.Set("field_type", _fieldType)
    return nil
}

// FieldType Getter
func (r AlibabaSeakingTranslateAPIRequest) GetFieldType() string {
    return r._fieldType
}
// Extra Setter
// 扩展信息
func (r *AlibabaSeakingTranslateAPIRequest) SetExtra(_extra *Extra) error {
    r._extra = _extra
    r.Set("extra", _extra)
    return nil
}

// Extra Getter
func (r AlibabaSeakingTranslateAPIRequest) GetExtra() *Extra {
    return r._extra
}
