package nlp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaonlpwordAPIRequest 文本语言词法分析 API请求
// taobao.nlp.word
//
// 提供文本语言处理中的词法分析功能,开放中文分词和词权重计算功能。
type TaobaonlpwordAPIRequest struct {
	model.Params
	// 功能类型选择：1)wType=1时提供分词功能，type=0时为基本粒度，type=1时为混合粒度，type=3时为基本粒度和混合粒度共同输出；
	_wType int64
	// 文本内容
	_text *Text
}

// NewTaobaonlpwordRequest 初始化TaobaonlpwordAPIRequest对象
func NewTaobaonlpwordRequest() *TaobaonlpwordAPIRequest {
	return &TaobaonlpwordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaonlpwordAPIRequest) GetApiMethodName() string {
	return "taobao.nlp.word"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaonlpwordAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaonlpwordAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWType is WType Setter
// 功能类型选择：1)wType=1时提供分词功能，type=0时为基本粒度，type=1时为混合粒度，type=3时为基本粒度和混合粒度共同输出；
func (r *TaobaonlpwordAPIRequest) SetWType(_wType int64) error {
	r._wType = _wType
	r.Set("w_type", _wType)
	return nil
}

// GetWType WType Getter
func (r TaobaonlpwordAPIRequest) GetWType() int64 {
	return r._wType
}

// SetText is Text Setter
// 文本内容
func (r *TaobaonlpwordAPIRequest) SetText(_text *Text) error {
	r._text = _text
	r.Set("text", _text)
	return nil
}

// GetText Text Getter
func (r TaobaonlpwordAPIRequest) GetText() *Text {
	return r._text
}
