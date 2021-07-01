package nlp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoNlpWordAPIRequest
文本语言词法分析 API请求
taobao.nlp.word

提供文本语言处理中的词法分析功能,开放中文分词和词权重计算功能。 */
type TaobaoNlpWordAPIRequest struct {
	model.Params
	// 功能类型选择：1)wType=1时提供分词功能，type=0时为基本粒度，type=1时为混合粒度，type=3时为基本粒度和混合粒度共同输出；
	_wType int64
	// 文本内容
	_text *Text
}

// New
