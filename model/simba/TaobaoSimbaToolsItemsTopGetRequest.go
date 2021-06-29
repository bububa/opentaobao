package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个关键词的推广组排名列表 API请求
taobao.simba.tools.items.top.get

取得一个关键词的推广组排名列表
*/
type TaobaoSimbaToolsItemsTopGetRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 关键词
    _keyword   string
    // 输入的必须是一个符合ipv4或者ipv6格式的IP地址
    _ip   string
}

// 初始化TaobaoSimbaToolsItemsTopGetRequest对象
func NewTaobaoSimbaToolsItemsTopGetRequest() *TaobaoSimbaToolsItemsTopGetRequest{
    return &TaobaoSimbaToolsItemsTopGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaToolsItemsTopGetRequest) GetApiMethodName() string {
    return "taobao.simba.tools.items.top.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaToolsItemsTopGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaToolsItemsTopGetRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaToolsItemsTopGetRequest) GetNick() string {
    return r._nick
}
// Keyword Setter
// 关键词
func (r *TaobaoSimbaToolsItemsTopGetRequest) SetKeyword(_keyword string) error {
    r._keyword = _keyword
    r.Set("keyword", _keyword)
    return nil
}

// Keyword Getter
func (r TaobaoSimbaToolsItemsTopGetRequest) GetKeyword() string {
    return r._keyword
}
// Ip Setter
// 输入的必须是一个符合ipv4或者ipv6格式的IP地址
func (r *TaobaoSimbaToolsItemsTopGetRequest) SetIp(_ip string) error {
    r._ip = _ip
    r.Set("ip", _ip)
    return nil
}

// Ip Getter
func (r TaobaoSimbaToolsItemsTopGetRequest) GetIp() string {
    return r._ip
}
