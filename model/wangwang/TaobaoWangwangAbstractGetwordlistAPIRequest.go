package wangwang

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词列表 API请求
taobao.wangwang.abstract.getwordlist

获取关键词列表，只支持json返回
*/
type TaobaoWangwangAbstractGetwordlistAPIRequest struct {
    model.Params
    // 传入参数的字符集
    _charset   string
}

// 初始化TaobaoWangwangAbstractGetwordlistAPIRequest对象
func NewTaobaoWangwangAbstractGetwordlistRequest() *TaobaoWangwangAbstractGetwordlistAPIRequest{
    return &TaobaoWangwangAbstractGetwordlistAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWangwangAbstractGetwordlistAPIRequest) GetApiMethodName() string {
    return "taobao.wangwang.abstract.getwordlist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWangwangAbstractGetwordlistAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Charset Setter
// 传入参数的字符集
func (r *TaobaoWangwangAbstractGetwordlistAPIRequest) SetCharset(_charset string) error {
    r._charset = _charset
    r.Set("charset", _charset)
    return nil
}

// Charset Getter
func (r TaobaoWangwangAbstractGetwordlistAPIRequest) GetCharset() string {
    return r._charset
}
