package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔短链信息查询 APIRequest
taobao.jst.sms.message.shorturl.query

聚石塔短链信息查询
*/
type TaobaoJstSmsMessageShorturlQueryRequest struct {
    model.Params

    // 短链名，即域名后的字符串
    shortName   string 

}

func NewTaobaoJstSmsMessageShorturlQueryRequest() *TaobaoJstSmsMessageShorturlQueryRequest{
    return &TaobaoJstSmsMessageShorturlQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstSmsMessageShorturlQueryRequest) GetApiMethodName() string {
    return "taobao.jst.sms.message.shorturl.query"
}

func (r TaobaoJstSmsMessageShorturlQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstSmsMessageShorturlQueryRequest) SetShortName(shortName string) error {
    r.shortName = shortName
    r.Set("short_name", shortName)
    return nil
}

func (r TaobaoJstSmsMessageShorturlQueryRequest) GetShortName() string {
    return r.shortName
}

