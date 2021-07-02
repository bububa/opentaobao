package wangwang

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWangwangAbstractDeletewordAPIRequest 删除关键词 API请求
// taobao.wangwang.abstract.deleteword
//
// 删除关键词，只支持json返回
type TaobaoWangwangAbstractDeletewordAPIRequest struct {
	model.Params
	// 关键词，长度大于0
	_word string
	// 传入参数的字符集
	_charset string
}

// NewTaobaoWangwangAbstractDeletewordRequest 初始化TaobaoWangwangAbstractDeletewordAPIRequest对象
func NewTaobaoWangwangAbstractDeletewordRequest() *TaobaoWangwangAbstractDeletewordAPIRequest {
	return &TaobaoWangwangAbstractDeletewordAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWangwangAbstractDeletewordAPIRequest) GetApiMethodName() string {
	return "taobao.wangwang.abstract.deleteword"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWangwangAbstractDeletewordAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Word Setter
// 关键词，长度大于0
func (r *TaobaoWangwangAbstractDeletewordAPIRequest) SetWord(_word string) error {
	r._word = _word
	r.Set("word", _word)
	return nil
}

// Get Word Getter
func (r TaobaoWangwangAbstractDeletewordAPIRequest) GetWord() string {
	return r._word
}

// Set is Charset Setter
// 传入参数的字符集
func (r *TaobaoWangwangAbstractDeletewordAPIRequest) SetCharset(_charset string) error {
	r._charset = _charset
	r.Set("charset", _charset)
	return nil
}

// Get Charset Getter
func (r TaobaoWangwangAbstractDeletewordAPIRequest) GetCharset() string {
	return r._charset
}
