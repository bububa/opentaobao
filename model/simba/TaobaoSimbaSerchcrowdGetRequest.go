package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据推广单元id获取搜索溢价人群 APIRequest
taobao.simba.serchcrowd.get

根据推广单元id获取搜索溢价人群
*/
type TaobaoSimbaSerchcrowdGetRequest struct {
    model.Params

    // 被操作者的淘宝昵称
    nick   string 

    // 推广单元id
    adgroupId   int64 

}

func NewTaobaoSimbaSerchcrowdGetRequest() *TaobaoSimbaSerchcrowdGetRequest{
    return &TaobaoSimbaSerchcrowdGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaSerchcrowdGetRequest) GetApiMethodName() string {
    return "taobao.simba.serchcrowd.get"
}

func (r TaobaoSimbaSerchcrowdGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaSerchcrowdGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaSerchcrowdGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaSerchcrowdGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaSerchcrowdGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

