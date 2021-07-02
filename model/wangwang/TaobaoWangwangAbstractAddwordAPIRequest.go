package wangwang

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWangwangAbstractAddwordAPIRequest 增加关键词 API请求
// taobao.wangwang.abstract.addword
//
// 增加关键词，只支持json返回
type TaobaoWangwangAbstractAddwordAPIRequest struct {
	model.Params
	// 关键词，长度大于0
	_word string
	// 传入参数的字符集
	_charset string
}

// NewTaobaoWangwangAbstractAddwordRequest 初始化TaobaoWangwangAbstractAddwordAPIRequest对象
func NewTaobaoWangwangAbstractAddwordRequest() *TaobaoWangwangAbstractAddwordAPIRequest {
	return &TaobaoWangwangAbstractAddwordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWangwangAbstractAddwordAPIRequest) GetApiMethodName() string {
	return "taobao.wangwang.abstract.addword"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWangwangAbstractAddwordAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetWord is Word Setter
// 关键词，长度大于0
func (r *TaobaoWangwangAbstractAddwordAPIRequest) SetWord(_word string) error {
	r._word = _word
	r.Set("word", _word)
	return nil
}

// GetWord Word Getter
func (r TaobaoWangwangAbstractAddwordAPIRequest) GetWord() string {
	return r._word
}

// SetCharset is Charset Setter
// 传入参数的字符集
func (r *TaobaoWangwangAbstractAddwordAPIRequest) SetCharset(_charset string) error {
	r._charset = _charset
	r.Set("charset", _charset)
	return nil
}

// GetCharset Charset Getter
func (r TaobaoWangwangAbstractAddwordAPIRequest) GetCharset() string {
	return r._charset
}
