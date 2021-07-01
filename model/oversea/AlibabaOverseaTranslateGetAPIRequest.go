package oversea

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaOverseaTranslateGetAPIRequest
获取文本翻译信息 API请求
alibaba.oversea.translate.get

根据传入的文本信息，获取其目标语言的翻译结果 */
type AlibabaOverseaTranslateGetAPIRequest struct {
	model.Params
	// 待翻译文本
	_text string
	// 源语种英文
	_sourceLang string
	// 目标语种中文
	_targetLang string
}

// NewAlibabaOverseaTranslateGetRequest 初始化AlibabaOverseaTranslateGetAPIRequest对象
func NewAlibabaOverseaTranslateGetRequest() *AlibabaOverseaTranslateGetAPIRequest {
	return &AlibabaOverseaTranslateGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOverseaTranslateGetAPIRequest) GetApiMethodName() string {
	return "alibaba.oversea.translate.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOverseaTranslateGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Text Setter
// 待翻译文本
func (r *AlibabaOverseaTranslateGetAPIRequest) SetText(_text string) error {
	r._text = _text
	r.Set("text", _text)
	return nil
}

// Get Text Getter
func (r AlibabaOverseaTranslateGetAPIRequest) GetText() string {
	return r._text
}

// Set is SourceLang Setter
// 源语种英文
func (r *AlibabaOverseaTranslateGetAPIRequest) SetSourceLang(_sourceLang string) error {
	r._sourceLang = _sourceLang
	r.Set("source_lang", _sourceLang)
	return nil
}

// Get SourceLang Getter
func (r AlibabaOverseaTranslateGetAPIRequest) GetSourceLang() string {
	return r._sourceLang
}

// Set is TargetLang Setter
// 目标语种中文
func (r *AlibabaOverseaTranslateGetAPIRequest) SetTargetLang(_targetLang string) error {
	r._targetLang = _targetLang
	r.Set("target_lang", _targetLang)
	return nil
}

// Get TargetLang Getter
func (r AlibabaOverseaTranslateGetAPIRequest) GetTargetLang() string {
	return r._targetLang
}
