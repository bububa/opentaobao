package oversea

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取文本翻译信息 APIRequest
alibaba.oversea.translate.get

根据传入的文本信息，获取其目标语言的翻译结果
*/
type AlibabaOverseaTranslateGetRequest struct {
    model.Params

    // 待翻译文本
    text   string 

    // 源语种英文
    sourceLang   string 

    // 目标语种中文
    targetLang   string 

}

func NewAlibabaOverseaTranslateGetRequest() *AlibabaOverseaTranslateGetRequest{
    return &AlibabaOverseaTranslateGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaOverseaTranslateGetRequest) GetApiMethodName() string {
    return "alibaba.oversea.translate.get"
}

func (r AlibabaOverseaTranslateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaOverseaTranslateGetRequest) SetText(text string) error {
    r.text = text
    r.Set("text", text)
    return nil
}

func (r AlibabaOverseaTranslateGetRequest) GetText() string {
    return r.text
}

func (r *AlibabaOverseaTranslateGetRequest) SetSourceLang(sourceLang string) error {
    r.sourceLang = sourceLang
    r.Set("source_lang", sourceLang)
    return nil
}

func (r AlibabaOverseaTranslateGetRequest) GetSourceLang() string {
    return r.sourceLang
}

func (r *AlibabaOverseaTranslateGetRequest) SetTargetLang(targetLang string) error {
    r.targetLang = targetLang
    r.Set("target_lang", targetLang)
    return nil
}

func (r AlibabaOverseaTranslateGetRequest) GetTargetLang() string {
    return r.targetLang
}

