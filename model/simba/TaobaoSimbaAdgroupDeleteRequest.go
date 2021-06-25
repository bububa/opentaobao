package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除一个推广组 APIRequest
taobao.simba.adgroup.delete

删除一个推广组
*/
type TaobaoSimbaAdgroupDeleteRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 推广组Id
    adgroupId   int64 

}

func NewTaobaoSimbaAdgroupDeleteRequest() *TaobaoSimbaAdgroupDeleteRequest{
    return &TaobaoSimbaAdgroupDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaAdgroupDeleteRequest) GetApiMethodName() string {
    return "taobao.simba.adgroup.delete"
}

func (r TaobaoSimbaAdgroupDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaAdgroupDeleteRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaAdgroupDeleteRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaAdgroupDeleteRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaAdgroupDeleteRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

