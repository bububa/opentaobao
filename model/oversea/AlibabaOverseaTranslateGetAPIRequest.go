package oversea

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOverseaTranslateGetAPIRequest 获取文本翻译信息 API请求
// alibaba.oversea.translate.get
//
// 根据传入的文本信息，获取其目标语言的翻译结果
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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaOverseaTranslateGetAPIRequest) Reset() {
	r._text = ""
	r._sourceLang = ""
	r._targetLang = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaOverseaTranslateGetAPIRequest) GetApiMethodName() string {
	return "alibaba.oversea.translate.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaOverseaTranslateGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaOverseaTranslateGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetText is Text Setter
// 待翻译文本
func (r *AlibabaOverseaTranslateGetAPIRequest) SetText(_text string) error {
	r._text = _text
	r.Set("text", _text)
	return nil
}

// GetText Text Getter
func (r AlibabaOverseaTranslateGetAPIRequest) GetText() string {
	return r._text
}

// SetSourceLang is SourceLang Setter
// 源语种英文
func (r *AlibabaOverseaTranslateGetAPIRequest) SetSourceLang(_sourceLang string) error {
	r._sourceLang = _sourceLang
	r.Set("source_lang", _sourceLang)
	return nil
}

// GetSourceLang SourceLang Getter
func (r AlibabaOverseaTranslateGetAPIRequest) GetSourceLang() string {
	return r._sourceLang
}

// SetTargetLang is TargetLang Setter
// 目标语种中文
func (r *AlibabaOverseaTranslateGetAPIRequest) SetTargetLang(_targetLang string) error {
	r._targetLang = _targetLang
	r.Set("target_lang", _targetLang)
	return nil
}

// GetTargetLang TargetLang Getter
func (r AlibabaOverseaTranslateGetAPIRequest) GetTargetLang() string {
	return r._targetLang
}

var poolAlibabaOverseaTranslateGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaOverseaTranslateGetRequest()
	},
}

// GetAlibabaOverseaTranslateGetRequest 从 sync.Pool 获取 AlibabaOverseaTranslateGetAPIRequest
func GetAlibabaOverseaTranslateGetAPIRequest() *AlibabaOverseaTranslateGetAPIRequest {
	return poolAlibabaOverseaTranslateGetAPIRequest.Get().(*AlibabaOverseaTranslateGetAPIRequest)
}

// ReleaseAlibabaOverseaTranslateGetAPIRequest 将 AlibabaOverseaTranslateGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaOverseaTranslateGetAPIRequest(v *AlibabaOverseaTranslateGetAPIRequest) {
	v.Reset()
	poolAlibabaOverseaTranslateGetAPIRequest.Put(v)
}
