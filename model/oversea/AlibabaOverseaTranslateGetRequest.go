package oversea

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取文本翻译信息 API请求
alibaba.oversea.translate.get

根据传入的文本信息，获取其目标语言的翻译结果
*/
type AlibabaOverseaTranslateGetRequest struct {
    model.Params
    // 待翻译文本
    _text   string
    // 源语种英文
    _sourceLang   string
    // 目标语种中文
    _targetLang   string
}

// 初始化AlibabaOverseaTranslateGetRequest对象
func NewAlibabaOverseaTranslateGetRequest() *AlibabaOverseaTranslateGetRequest{
    return &AlibabaOverseaTranslateGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOverseaTranslateGetRequest) GetApiMethodName() string {
    return "alibaba.oversea.translate.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOverseaTranslateGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Text Setter
// 待翻译文本
func (r *AlibabaOverseaTranslateGetRequest) SetText(_text string) error {
    r._text = _text
    r.Set("text", _text)
    return nil
}

// Text Getter
func (r AlibabaOverseaTranslateGetRequest) GetText() string {
    return r._text
}
// SourceLang Setter
// 源语种英文
func (r *AlibabaOverseaTranslateGetRequest) SetSourceLang(_sourceLang string) error {
    r._sourceLang = _sourceLang
    r.Set("source_lang", _sourceLang)
    return nil
}

// SourceLang Getter
func (r AlibabaOverseaTranslateGetRequest) GetSourceLang() string {
    return r._sourceLang
}
// TargetLang Setter
// 目标语种中文
func (r *AlibabaOverseaTranslateGetRequest) SetTargetLang(_targetLang string) error {
    r._targetLang = _targetLang
    r.Set("target_lang", _targetLang)
    return nil
}

// TargetLang Getter
func (r AlibabaOverseaTranslateGetRequest) GetTargetLang() string {
    return r._targetLang
}
