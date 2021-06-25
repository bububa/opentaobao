package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广组的所有关键词 APIRequest
taobao.simba.keywordsbyadgroupid.get

取得一个推广组的所有关键词
*/
type TaobaoSimbaKeywordsbyadgroupidGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 推广组Id
    adgroupId   int64 

}

func NewTaobaoSimbaKeywordsbyadgroupidGetRequest() *TaobaoSimbaKeywordsbyadgroupidGetRequest{
    return &TaobaoSimbaKeywordsbyadgroupidGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaKeywordsbyadgroupidGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywordsbyadgroupid.get"
}

func (r TaobaoSimbaKeywordsbyadgroupidGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaKeywordsbyadgroupidGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaKeywordsbyadgroupidGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaKeywordsbyadgroupidGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaKeywordsbyadgroupidGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

