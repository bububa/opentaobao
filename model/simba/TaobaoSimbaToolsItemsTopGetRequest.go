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
    nick   string
    // 关键词
    keyword   string
    // 输入的必须是一个符合ipv4或者ipv6格式的IP地址
    ip   string
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
func (r *TaobaoSimbaToolsItemsTopGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaToolsItemsTopGetRequest) GetNick() string {
    return r.nick
}
// Keyword Setter
// 关键词
func (r *TaobaoSimbaToolsItemsTopGetRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

// Keyword Getter
func (r TaobaoSimbaToolsItemsTopGetRequest) GetKeyword() string {
    return r.keyword
}
// Ip Setter
// 输入的必须是一个符合ipv4或者ipv6格式的IP地址
func (r *TaobaoSimbaToolsItemsTopGetRequest) SetIp(ip string) error {
    r.ip = ip
    r.Set("ip", ip)
    return nil
}

// Ip Getter
func (r TaobaoSimbaToolsItemsTopGetRequest) GetIp() string {
    return r.ip
}
