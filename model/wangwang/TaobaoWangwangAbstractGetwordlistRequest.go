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
type TaobaoWangwangAbstractGetwordlistRequest struct {
    model.Params
    // 传入参数的字符集
    _charset   string
}

// 初始化TaobaoWangwangAbstractGetwordlistRequest对象
func NewTaobaoWangwangAbstractGetwordlistRequest() *TaobaoWangwangAbstractGetwordlistRequest{
    return &TaobaoWangwangAbstractGetwordlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoWangwangAbstractGetwordlistRequest) GetApiMethodName() string {
    return "taobao.wangwang.abstract.getwordlist"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoWangwangAbstractGetwordlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Charset Setter
// 传入参数的字符集
func (r *TaobaoWangwangAbstractGetwordlistRequest) SetCharset(_charset string) error {
    r._charset = _charset
    r.Set("charset", _charset)
    return nil
}

// Charset Getter
func (r TaobaoWangwangAbstractGetwordlistRequest) GetCharset() string {
    return r._charset
}
