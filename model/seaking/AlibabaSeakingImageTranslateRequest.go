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
    _identifier   string
    // 目标语种
    _targetLang   string
    // 源语种
    _sourceLang   string
    // 调用来源(erp名称)
    _identifierType   string
    // 原图url
    _url   string
    // 扩展信息
    _extra   *Extra
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
func (r *AlibabaSeakingImagetranslateRequest) SetIdentifier(_identifier string) error {
    r._identifier = _identifier
    r.Set("identifier", _identifier)
    return nil
}

// Identifier Getter
func (r AlibabaSeakingImagetranslateRequest) GetIdentifier() string {
    return r._identifier
}
// TargetLang Setter
// 目标语种
func (r *AlibabaSeakingImagetranslateRequest) SetTargetLang(_targetLang string) error {
    r._targetLang = _targetLang
    r.Set("target_lang", _targetLang)
    return nil
}

// TargetLang Getter
func (r AlibabaSeakingImagetranslateRequest) GetTargetLang() string {
    return r._targetLang
}
// SourceLang Setter
// 源语种
func (r *AlibabaSeakingImagetranslateRequest) SetSourceLang(_sourceLang string) error {
    r._sourceLang = _sourceLang
    r.Set("source_lang", _sourceLang)
    return nil
}

// SourceLang Getter
func (r AlibabaSeakingImagetranslateRequest) GetSourceLang() string {
    return r._sourceLang
}
// IdentifierType Setter
// 调用来源(erp名称)
func (r *AlibabaSeakingImagetranslateRequest) SetIdentifierType(_identifierType string) error {
    r._identifierType = _identifierType
    r.Set("identifier_type", _identifierType)
    return nil
}

// IdentifierType Getter
func (r AlibabaSeakingImagetranslateRequest) GetIdentifierType() string {
    return r._identifierType
}
// Url Setter
// 原图url
func (r *AlibabaSeakingImagetranslateRequest) SetUrl(_url string) error {
    r._url = _url
    r.Set("url", _url)
    return nil
}

// Url Getter
func (r AlibabaSeakingImagetranslateRequest) GetUrl() string {
    return r._url
}
// Extra Setter
// 扩展信息
func (r *AlibabaSeakingImagetranslateRequest) SetExtra(_extra *Extra) error {
    r._extra = _extra
    r.Set("extra", _extra)
    return nil
}

// Extra Getter
func (r AlibabaSeakingImagetranslateRequest) GetExtra() *Extra {
    return r._extra
}
