package wangwang

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除关键词 API请求
taobao.wangwang.abstract.deleteword

删除关键词，只支持json返回
*/
type TaobaoWangwangAbstractDeletewordRequest struct {
    model.Params
    // 关键词，长度大于0
    word   string
    // 传入参数的字符集
    charset   string
}

// 初始化TaobaoWangwangAbstractDeletewordRequest对象
func NewTaobaoWangwangAbstractDeletewordRequest() *TaobaoWangwangAbstractDeletewordRequest{
    return &TaobaoWangwangAbstractDeletewordRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWangwangAbstractDeletewordRequest) GetApiMethodName() string {
    return "taobao.wangwang.abstract.deleteword"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWangwangAbstractDeletewordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Word Setter
// 关键词，长度大于0
func (r *TaobaoWangwangAbstractDeletewordRequest) SetWord(word string) error {
    r.word = word
    r.Set("word", word)
    return nil
}

// Word Getter
func (r TaobaoWangwangAbstractDeletewordRequest) GetWord() string {
    return r.word
}
// Charset Setter
// 传入参数的字符集
func (r *TaobaoWangwangAbstractDeletewordRequest) SetCharset(charset string) error {
    r.charset = charset
    r.Set("charset", charset)
    return nil
}

// Charset Getter
func (r TaobaoWangwangAbstractDeletewordRequest) GetCharset() string {
    return r.charset
}
