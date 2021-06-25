package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
(新)销量明星删除推广单元接口 APIRequest
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

func NewTaobaoSimbaSalestarAdgroupDeleteRequest() *TaobaoSimbaSalestarAdgroupDeleteRequest{
    return &TaobaoSimbaSalestarAdgroupDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaSalestarAdgroupDeleteRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.adgroup.delete"
}

func (r TaobaoSimbaSalestarAdgroupDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaSalestarAdgroupDeleteRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaSalestarAdgroupDeleteRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaSalestarAdgroupDeleteRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaSalestarAdgroupDeleteRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

