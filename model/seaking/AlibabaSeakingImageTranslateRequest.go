package seaking

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图片机器翻译 APIRequest
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

func NewAlibabaSeakingImagetranslateRequest() *AlibabaSeakingImagetranslateRequest{
    return &AlibabaSeakingImagetranslateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSeakingImagetranslateRequest) GetApiMethodName() string {
    return "alibaba.seaking.imagetranslate"
}

func (r AlibabaSeakingImagetranslateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSeakingImagetranslateRequest) SetIdentifier(identifier string) error {
    r.identifier = identifier
    r.Set("identifier", identifier)
    return nil
}

func (r AlibabaSeakingImagetranslateRequest) GetIdentifier() string {
    return r.identifier
}

func (r *AlibabaSeakingImagetranslateRequest) SetTargetLang(targetLang string) error {
    r.targetLang = targetLang
    r.Set("target_lang", targetLang)
    return nil
}

func (r AlibabaSeakingImagetranslateRequest) GetTargetLang() string {
    return r.targetLang
}

func (r *AlibabaSeakingImagetranslateRequest) SetSourceLang(sourceLang string) error {
    r.sourceLang = sourceLang
    r.Set("source_lang", sourceLang)
    return nil
}

func (r AlibabaSeakingImagetranslateRequest) GetSourceLang() string {
    return r.sourceLang
}

func (r *AlibabaSeakingImagetranslateRequest) SetIdentifierType(identifierType string) error {
    r.identifierType = identifierType
    r.Set("identifier_type", identifierType)
    return nil
}

func (r AlibabaSeakingImagetranslateRequest) GetIdentifierType() string {
    return r.identifierType
}

func (r *AlibabaSeakingImagetranslateRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r AlibabaSeakingImagetranslateRequest) GetUrl() string {
    return r.url
}

func (r *AlibabaSeakingImagetranslateRequest) SetExtra(extra *Extra) error {
    r.extra = extra
    r.Set("extra", extra)
    return nil
}

func (r AlibabaSeakingImagetranslateRequest) GetExtra() *Extra {
    return r.extra
}

