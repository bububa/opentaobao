package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广组的所有关键词的质量得分或者根据关键词Id列表取得一组关键词的质量得分 APIRequest
taobao.simba.keywords.qscore.get

取得一个推广组的所有关键词的质量得分列表
*/
type TaobaoSimbaKeywordsQscoreGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 推广组Id
    adgroupId   int64 

}

func NewTaobaoSimbaKeywordsQscoreGetRequest() *TaobaoSimbaKeywordsQscoreGetRequest{
    return &TaobaoSimbaKeywordsQscoreGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaKeywordsQscoreGetRequest) GetApiMethodName() string {
    return "taobao.simba.keywords.qscore.get"
}

func (r TaobaoSimbaKeywordsQscoreGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaKeywordsQscoreGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaKeywordsQscoreGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaKeywordsQscoreGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaKeywordsQscoreGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

