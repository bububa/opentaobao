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
type TaobaoJstSmsMessageShorturlQueryAPIRequest struct {
    model.Params
    // 短链名，即域名后的字符串
    _shortName   string
}

// 初始化TaobaoJstSmsMessageShorturlQueryAPIRequest对象
func NewTaobaoJstSmsMessageShorturlQueryRequest() *TaobaoJstSmsMessageShorturlQueryAPIRequest{
    return &TaobaoJstSmsMessageShorturlQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsMessageShorturlQueryAPIRequest) GetApiMethodName() string {
    return "taobao.jst.sms.message.shorturl.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsMessageShorturlQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShortName Setter
// 短链名，即域名后的字符串
func (r *TaobaoJstSmsMessageShorturlQueryAPIRequest) SetShortName(_shortName string) error {
    r._shortName = _shortName
    r.Set("short_name", _shortName)
    return nil
}

// ShortName Getter
func (r TaobaoJstSmsMessageShorturlQueryAPIRequest) GetShortName() string {
    return r._shortName
}
