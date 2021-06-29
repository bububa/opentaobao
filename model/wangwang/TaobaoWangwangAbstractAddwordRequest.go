package wangwang

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加关键词 API请求
taobao.wangwang.abstract.addword

增加关键词，只支持json返回
*/
type TaobaoWangwangAbstractAddwordRequest struct {
    model.Params
    // 关键词，长度大于0
    _word   string
    // 传入参数的字符集
    _charset   string
}

// 初始化TaobaoWangwangAbstractAddwordRequest对象
func NewTaobaoWangwangAbstractAddwordRequest() *TaobaoWangwangAbstractAddwordRequest{
    return &TaobaoWangwangAbstractAddwordRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWangwangAbstractAddwordRequest) GetApiMethodName() string {
    return "taobao.wangwang.abstract.addword"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWangwangAbstractAddwordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Word Setter
// 关键词，长度大于0
func (r *TaobaoWangwangAbstractAddwordRequest) SetWord(_word string) error {
    r._word = _word
    r.Set("word", _word)
    return nil
}

// Word Getter
func (r TaobaoWangwangAbstractAddwordRequest) GetWord() string {
    return r._word
}
// Charset Setter
// 传入参数的字符集
func (r *TaobaoWangwangAbstractAddwordRequest) SetCharset(_charset string) error {
    r._charset = _charset
    r.Set("charset", _charset)
    return nil
}

// Charset Getter
func (r TaobaoWangwangAbstractAddwordRequest) GetCharset() string {
    return r._charset
}
