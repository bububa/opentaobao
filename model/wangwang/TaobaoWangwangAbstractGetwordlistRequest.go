package wangwang

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词列表 APIRequest
taobao.wangwang.abstract.getwordlist

获取关键词列表，只支持json返回
*/
type TaobaoWangwangAbstractGetwordlistRequest struct {
    model.Params

    // 传入参数的字符集
    charset   string 

}

func NewTaobaoWangwangAbstractGetwordlistRequest() *TaobaoWangwangAbstractGetwordlistRequest{
    return &TaobaoWangwangAbstractGetwordlistRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWangwangAbstractGetwordlistRequest) GetApiMethodName() string {
    return "taobao.wangwang.abstract.getwordlist"
}

func (r TaobaoWangwangAbstractGetwordlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWangwangAbstractGetwordlistRequest) SetCharset(charset string) error {
    r.charset = charset
    r.Set("charset", charset)
    return nil
}

func (r TaobaoWangwangAbstractGetwordlistRequest) GetCharset() string {
    return r.charset
}

