package nlp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoNlpPreprocessAPIRequest 文本语言预处理 API请求
// taobao.nlp.preprocess
//
// 实现文本语言处理中的预处理功能，如实现文字繁简转换、文字转拼音、文字拆分、谐音同音字替换和形似字替换。
type TaobaoNlpPreprocessAPIRequest struct {
	model.Params
	// 谐音字转换、形似字转换需提供关键词进行替换，关键词之间以#分隔。keyword示例：兼职#招聘#微信、天猫#日结#微信#招聘#加微
	_keyword string
	// 文本内容
	_text *Text
	// 1)繁简字转换：func_type=1，对应type =1 繁转简 type=2 简转繁；2)拆分字转换：func_type =2，对应type=1 文字拆分 type=2 拆分字合并；3)文字转拼音：func_type =3，对应type=1 文字转拼音 type=2 拼音+声调；4)谐音同音字替换：func_type =4，对应type=1 谐音字替换 type=2 同音字替换；5)形似字替换：func_type =5，对应type=1 形似字替换;
	_funcType int64
}

// NewTaobaoNlpPreprocessRequest 初始化TaobaoNlpPreprocessAPIRequest对象
func NewTaobaoNlpPreprocessRequest() *TaobaoNlpPreprocessAPIRequest {
	return &TaobaoNlpPreprocessAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoNlpPreprocessAPIRequest) Reset() {
	r._keyword = ""
	r._text = nil
	r._funcType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoNlpPreprocessAPIRequest) GetApiMethodName() string {
	return "taobao.nlp.preprocess"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoNlpPreprocessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoNlpPreprocessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 谐音字转换、形似字转换需提供关键词进行替换，关键词之间以#分隔。keyword示例：兼职#招聘#微信、天猫#日结#微信#招聘#加微
func (r *TaobaoNlpPreprocessAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r TaobaoNlpPreprocessAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetText is Text Setter
// 文本内容
func (r *TaobaoNlpPreprocessAPIRequest) SetText(_text *Text) error {
	r._text = _text
	r.Set("text", _text)
	return nil
}

// GetText Text Getter
func (r TaobaoNlpPreprocessAPIRequest) GetText() *Text {
	return r._text
}

// SetFuncType is FuncType Setter
// 1)繁简字转换：func_type=1，对应type =1 繁转简 type=2 简转繁；2)拆分字转换：func_type =2，对应type=1 文字拆分 type=2 拆分字合并；3)文字转拼音：func_type =3，对应type=1 文字转拼音 type=2 拼音+声调；4)谐音同音字替换：func_type =4，对应type=1 谐音字替换 type=2 同音字替换；5)形似字替换：func_type =5，对应type=1 形似字替换;
func (r *TaobaoNlpPreprocessAPIRequest) SetFuncType(_funcType int64) error {
	r._funcType = _funcType
	r.Set("func_type", _funcType)
	return nil
}

// GetFuncType FuncType Getter
func (r TaobaoNlpPreprocessAPIRequest) GetFuncType() int64 {
	return r._funcType
}

var poolTaobaoNlpPreprocessAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoNlpPreprocessRequest()
	},
}

// GetTaobaoNlpPreprocessRequest 从 sync.Pool 获取 TaobaoNlpPreprocessAPIRequest
func GetTaobaoNlpPreprocessAPIRequest() *TaobaoNlpPreprocessAPIRequest {
	return poolTaobaoNlpPreprocessAPIRequest.Get().(*TaobaoNlpPreprocessAPIRequest)
}

// ReleaseTaobaoNlpPreprocessAPIRequest 将 TaobaoNlpPreprocessAPIRequest 放入 sync.Pool
func ReleaseTaobaoNlpPreprocessAPIRequest(v *TaobaoNlpPreprocessAPIRequest) {
	v.Reset()
	poolTaobaoNlpPreprocessAPIRequest.Put(v)
}
