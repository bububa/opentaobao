package oversea

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaoverseatranslategetAPIRequest 获取文本翻译信息 API请求
// alibaba.oversea.translate.get
//
// 根据传入的文本信息，获取其目标语言的翻译结果
type AlibabaoverseatranslategetAPIRequest struct {
	model.Params
	// 待翻译文本
	_text string
	// 源语种英文
	_sourceLang string
	// 目标语种中文
	_targetLang string
}

// NewAlibabaoverseatranslategetRequest 初始化AlibabaoverseatranslategetAPIRequest对象
func NewAlibabaoverseatranslategetRequest() *AlibabaoverseatranslategetAPIRequest {
	return &AlibabaoverseatranslategetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaoverseatranslategetAPIRequest) GetApiMethodName() string {
	return "alibaba.oversea.translate.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaoverseatranslategetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaoverseatranslategetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetText is Text Setter
// 待翻译文本
func (r *AlibabaoverseatranslategetAPIRequest) SetText(_text string) error {
	r._text = _text
	r.Set("text", _text)
	return nil
}

// GetText Text Getter
func (r AlibabaoverseatranslategetAPIRequest) GetText() string {
	return r._text
}

// SetSourceLang is SourceLang Setter
// 源语种英文
func (r *AlibabaoverseatranslategetAPIRequest) SetSourceLang(_sourceLang string) error {
	r._sourceLang = _sourceLang
	r.Set("source_lang", _sourceLang)
	return nil
}

// GetSourceLang SourceLang Getter
func (r AlibabaoverseatranslategetAPIRequest) GetSourceLang() string {
	return r._sourceLang
}

// SetTargetLang is TargetLang Setter
// 目标语种中文
func (r *AlibabaoverseatranslategetAPIRequest) SetTargetLang(_targetLang string) error {
	r._targetLang = _targetLang
	r.Set("target_lang", _targetLang)
	return nil
}

// GetTargetLang TargetLang Getter
func (r AlibabaoverseatranslategetAPIRequest) GetTargetLang() string {
	return r._targetLang
}
