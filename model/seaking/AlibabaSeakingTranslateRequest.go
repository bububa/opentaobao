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
type AlibabaSeakingTranslateRequest struct {
    model.Params
    // 定制用户id
    identifier   string
    // 目标语种
    targetLang   string
    // 源语种
    sourceLang   string
    // 原文
    sourceText   string
    // 原文格式(text/html)
    sourceFormat   string
    // 定制用户类型
    identifierType   string
    // 原文类型(title: 标题/offer: 详描/message: 消息)
    fieldType   string
    // 扩展信息
    extra   *Extra
}

// 初始化AlibabaSeakingTranslateRequest对象
func NewAlibabaSeakingTranslateRequest() *AlibabaSeakingTranslateRequest{
    return &AlibabaSeakingTranslateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingTranslateRequest) GetApiMethodName() string {
    return "alibaba.seaking.translate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingTranslateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Identifier Setter
// 定制用户id
func (r *AlibabaSeakingTranslateRequest) SetIdentifier(identifier string) error {
    r.identifier = identifier
    r.Set("identifier", identifier)
    return nil
}

// Identifier Getter
func (r AlibabaSeakingTranslateRequest) GetIdentifier() string {
    return r.identifier
}
// TargetLang Setter
// 目标语种
func (r *AlibabaSeakingTranslateRequest) SetTargetLang(targetLang string) error {
    r.targetLang = targetLang
    r.Set("target_lang", targetLang)
    return nil
}

// TargetLang Getter
func (r AlibabaSeakingTranslateRequest) GetTargetLang() string {
    return r.targetLang
}
// SourceLang Setter
// 源语种
func (r *AlibabaSeakingTranslateRequest) SetSourceLang(sourceLang string) error {
    r.sourceLang = sourceLang
    r.Set("source_lang", sourceLang)
    return nil
}

// SourceLang Getter
func (r AlibabaSeakingTranslateRequest) GetSourceLang() string {
    return r.sourceLang
}
// SourceText Setter
// 原文
func (r *AlibabaSeakingTranslateRequest) SetSourceText(sourceText string) error {
    r.sourceText = sourceText
    r.Set("source_text", sourceText)
    return nil
}

// SourceText Getter
func (r AlibabaSeakingTranslateRequest) GetSourceText() string {
    return r.sourceText
}
// SourceFormat Setter
// 原文格式(text/html)
func (r *AlibabaSeakingTranslateRequest) SetSourceFormat(sourceFormat string) error {
    r.sourceFormat = sourceFormat
    r.Set("source_format", sourceFormat)
    return nil
}

// SourceFormat Getter
func (r AlibabaSeakingTranslateRequest) GetSourceFormat() string {
    return r.sourceFormat
}
// IdentifierType Setter
// 定制用户类型
func (r *AlibabaSeakingTranslateRequest) SetIdentifierType(identifierType string) error {
    r.identifierType = identifierType
    r.Set("identifier_type", identifierType)
    return nil
}

// IdentifierType Getter
func (r AlibabaSeakingTranslateRequest) GetIdentifierType() string {
    return r.identifierType
}
// FieldType Setter
// 原文类型(title: 标题/offer: 详描/message: 消息)
func (r *AlibabaSeakingTranslateRequest) SetFieldType(fieldType string) error {
    r.fieldType = fieldType
    r.Set("field_type", fieldType)
    return nil
}

// FieldType Getter
func (r AlibabaSeakingTranslateRequest) GetFieldType() string {
    return r.fieldType
}
// Extra Setter
// 扩展信息
func (r *AlibabaSeakingTranslateRequest) SetExtra(extra *Extra) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

// Extra Getter
func (r AlibabaSeakingTranslateRequest) GetExtra() *Extra {
    return r.extra
}
