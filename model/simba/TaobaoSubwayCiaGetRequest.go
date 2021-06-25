package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询单元智能出价信息 APIRequest
taobao.subway.cia.get

查询单元智能出价信息
*/
type TaobaoSubwayCiaGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 推广组Id
    adgroupId   int64 

}

func NewTaobaoSubwayCiaGetRequest() *TaobaoSubwayCiaGetRequest{
    return &TaobaoSubwayCiaGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSubwayCiaGetRequest) GetApiMethodName() string {
    return "taobao.subway.cia.get"
}

func (r TaobaoSubwayCiaGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSubwayCiaGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSubwayCiaGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSubwayCiaGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSubwayCiaGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

