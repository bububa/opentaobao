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

// New
