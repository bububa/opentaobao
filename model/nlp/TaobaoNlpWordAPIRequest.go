package nlp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoNlpWordAPIRequest 文本语言词法分析 API请求
// taobao.nlp.word
//
// 提供文本语言处理中的词法分析功能,开放中文分词和词权重计算功能。
type TaobaoNlpWordAPIRequest struct {
	model.Params
	// 功能类型选择：1)wType=1时提供分词功能，type=0时为基本粒度，type=1时为混合粒度，type=3时为基本粒度和混合粒度共同输出；
	_wType int64
	// 文本内容
	_text *Text
}

// NewTaobaoNlpWordRequest 初始化TaobaoNlpWordAPIRequest对象
func NewTaobaoNlpWordRequest() *TaobaoNlpWordAPIRequest {
	return &TaobaoNlpWordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoNlpWordAPIRequest) GetApiMethodName() string {
	return "taobao.nlp.word"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoNlpWordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoNlpWordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWType is WType Setter
// 功能类型选择：1)wType=1时提供分词功能，type=0时为基本粒度，type=1时为混合粒度，type=3时为基本粒度和混合粒度共同输出；
func (r *TaobaoNlpWordAPIRequest) SetWType(_wType int64) error {
	r._wType = _wType
	r.Set("w_type", _wType)
	return nil
}

// GetWType WType Getter
func (r TaobaoNlpWordAPIRequest) GetWType() int64 {
	return r._wType
}

// SetText is Text Setter
// 文本内容
func (r *TaobaoNlpWordAPIRequest) SetText(_text *Text) error {
	r._text = _text
	r.Set("text", _text)
	return nil
}

// GetText Text Getter
func (r TaobaoNlpWordAPIRequest) GetText() *Text {
	return r._text
}
