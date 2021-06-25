package wangwang

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
模糊查询服务初始化 APIRequest
taobao.wangwang.abstract.initialize

模糊查询服务初始化，只支持json返回
*/
type TaobaoWangwangAbstractInitializeRequest struct {
    model.Params

    // 传入参数的字符集
    charset   string 

}

func NewTaobaoWangwangAbstractInitializeRequest() *TaobaoWangwangAbstractInitializeRequest{
    return &TaobaoWangwangAbstractInitializeRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWangwangAbstractInitializeRequest) GetApiMethodName() string {
    return "taobao.wangwang.abstract.initialize"
}

func (r TaobaoWangwangAbstractInitializeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWangwangAbstractInitializeRequest) SetCharset(charset string) error {
    r.charset = charset
    r.Set("charset", charset)
    return nil
}

func (r TaobaoWangwangAbstractInitializeRequest) GetCharset() string {
    return r.charset
}

