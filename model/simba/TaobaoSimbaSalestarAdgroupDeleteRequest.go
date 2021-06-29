package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
(新)销量明星删除推广单元接口 API请求
taobao.simba.salestar.adgroup.delete

删除一个推广组
*/
type TaobaoSimbaSalestarAdgroupDeleteRequest struct {
    model.Params
    // 主人昵称
    nick   string
    // 推广组Id
    adgroupId   int64
}

// 初始化TaobaoSimbaSalestarAdgroupDeleteRequest对象
func NewTaobaoSimbaSalestarAdgroupDeleteRequest() *TaobaoSimbaSalestarAdgroupDeleteRequest{
    return &TaobaoSimbaSalestarAdgroupDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarAdgroupDeleteRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.adgroup.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarAdgroupDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaSalestarAdgroupDeleteRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaSalestarAdgroupDeleteRequest) GetNick() string {
    return r.nick
}
// AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaSalestarAdgroupDeleteRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaSalestarAdgroupDeleteRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
