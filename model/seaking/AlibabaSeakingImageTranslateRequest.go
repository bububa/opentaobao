package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片机器翻译 API请求
alibaba.seaking.imagetranslate

图片机器翻译
*/
type AlibabaSeakingImagetranslateRequest struct {
    model.Params
    // erp用户id
    identifier   string
    // 目标语种
    targetLang   string
    // 源语种
    sourceLang   string
    // 调用来源(erp名称)
    identifierType   string
    // 原图url
    url   string
    // 扩展信息
    extra   *Extra
}

// 初始化AlibabaSeakingImagetranslateRequest对象
func NewAlibabaSeakingImagetranslateRequest() *AlibabaSeakingImagetranslateRequest{
    return &AlibabaSeakingImagetranslateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSeakingImagetranslateRequest) GetApiMethodName() string {
    return "alibaba.seaking.imagetranslate"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSeakingImagetranslateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Identifier Setter
// erp用户id
func (r *AlibabaSeakingImagetranslateRequest) SetIdentifier(identifier string) error {
    r.identifier = identifier
    r.Set("identifier", identifier)
    return nil
}

// Identifier Getter
func (r AlibabaSeakingImagetranslateRequest) GetIdentifier() string {
    return r.identifier
}
// TargetLang Setter
// 目标语种
func (r *AlibabaSeakingImagetranslateRequest) SetTargetLang(targetLang string) error {
    r.targetLang = targetLang
    r.Set("target_lang", targetLang)
    return nil
}

// TargetLang Getter
func (r AlibabaSeakingImagetranslateRequest) GetTargetLang() string {
    return r.targetLang
}
// SourceLang Setter
// 源语种
func (r *AlibabaSeakingImagetranslateRequest) SetSourceLang(sourceLang string) error {
    r.sourceLang = sourceLang
    r.Set("source_lang", sourceLang)
    return nil
}

// SourceLang Getter
func (r AlibabaSeakingImagetranslateRequest) GetSourceLang() string {
    return r.sourceLang
}
// IdentifierType Setter
// 调用来源(erp名称)
func (r *AlibabaSeakingImagetranslateRequest) SetIdentifierType(identifierType string) error {
    r.identifierType = identifierType
    r.Set("identifier_type", identifierType)
    return nil
}

// IdentifierType Getter
func (r AlibabaSeakingImagetranslateRequest) GetIdentifierType() string {
    return r.identifierType
}
// Url Setter
// 原图url
func (r *AlibabaSeakingImagetranslateRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

// Url Getter
func (r AlibabaSeakingImagetranslateRequest) GetUrl() string {
    return r.url
}
// Extra Setter
// 扩展信息
func (r *AlibabaSeakingImagetranslateRequest) SetExtra(extra *Extra) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

// Extra Getter
func (r AlibabaSeakingImagetranslateRequest) GetExtra() *Extra {
    return r.extra
}
