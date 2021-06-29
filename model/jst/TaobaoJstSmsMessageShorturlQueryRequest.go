package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔短链信息查询 API请求
taobao.jst.sms.message.shorturl.query

聚石塔短链信息查询
*/
type TaobaoJstSmsMessageShorturlQueryRequest struct {
    model.Params
    // 短链名，即域名后的字符串
    shortName   string
}

// 初始化TaobaoJstSmsMessageShorturlQueryRequest对象
func NewTaobaoJstSmsMessageShorturlQueryRequest() *TaobaoJstSmsMessageShorturlQueryRequest{
    return &TaobaoJstSmsMessageShorturlQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsMessageShorturlQueryRequest) GetApiMethodName() string {
    return "taobao.jst.sms.message.shorturl.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsMessageShorturlQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShortName Setter
// 短链名，即域名后的字符串
func (r *TaobaoJstSmsMessageShorturlQueryRequest) SetShortName(shortName string) error {
    r.shortName = shortName
    r.Set("short_name", shortName)
    return nil
}

// ShortName Getter
func (r TaobaoJstSmsMessageShorturlQueryRequest) GetShortName() string {
    return r.shortName
}
