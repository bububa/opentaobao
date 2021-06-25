package wangwang

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
增加关键词 APIRequest
taobao.wangwang.abstract.addword

增加关键词，只支持json返回
*/
type TaobaoWangwangAbstractAddwordRequest struct {
    model.Params

    // 关键词，长度大于0
    word   string 

    // 传入参数的字符集
    charset   string 

}

func NewTaobaoWangwangAbstractAddwordRequest() *TaobaoWangwangAbstractAddwordRequest{
    return &TaobaoWangwangAbstractAddwordRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWangwangAbstractAddwordRequest) GetApiMethodName() string {
    return "taobao.wangwang.abstract.addword"
}

func (r TaobaoWangwangAbstractAddwordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWangwangAbstractAddwordRequest) SetWord(word string) error {
    r.word = word
    r.Set("word", word)
    return nil
}

func (r TaobaoWangwangAbstractAddwordRequest) GetWord() string {
    return r.word
}

func (r *TaobaoWangwangAbstractAddwordRequest) SetCharset(charset string) error {
    r.charset = charset
    r.Set("charset", charset)
    return nil
}

func (r TaobaoWangwangAbstractAddwordRequest) GetCharset() string {
    return r.charset
}

