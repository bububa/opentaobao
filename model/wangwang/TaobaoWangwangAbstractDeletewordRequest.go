package wangwang

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除关键词 APIRequest
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

func NewTaobaoWangwangAbstractDeletewordRequest() *TaobaoWangwangAbstractDeletewordRequest{
    return &TaobaoWangwangAbstractDeletewordRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWangwangAbstractDeletewordRequest) GetApiMethodName() string {
    return "taobao.wangwang.abstract.deleteword"
}

func (r TaobaoWangwangAbstractDeletewordRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWangwangAbstractDeletewordRequest) SetWord(word string) error {
    r.word = word
    r.Set("word", word)
    return nil
}

func (r TaobaoWangwangAbstractDeletewordRequest) GetWord() string {
    return r.word
}

func (r *TaobaoWangwangAbstractDeletewordRequest) SetCharset(charset string) error {
    r.charset = charset
    r.Set("charset", charset)
    return nil
}

func (r TaobaoWangwangAbstractDeletewordRequest) GetCharset() string {
    return r.charset
}

