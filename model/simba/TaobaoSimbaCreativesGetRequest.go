package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获得创意 APIRequest
taobao.simba.creatives.get

取得一个推广组的所有创意或者根据一个创意Id列表取得一组创意；<br/>如果同时提供了推广组Id和创意id列表，则优先使用推广组Id；
*/
type TaobaoSimbaCreativesGetRequest struct {
    model.Params

    // 主人昵称
    nick   string 

    // 创意Id数组，最多200个
    creativeIds   []int64 

    // 推广组Id
    adgroupId   int64 

}

func NewTaobaoSimbaCreativesGetRequest() *TaobaoSimbaCreativesGetRequest{
    return &TaobaoSimbaCreativesGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaCreativesGetRequest) GetApiMethodName() string {
    return "taobao.simba.creatives.get"
}

func (r TaobaoSimbaCreativesGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaCreativesGetRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaCreativesGetRequest) GetNick() string {
    return r.nick
}

func (r *TaobaoSimbaCreativesGetRequest) SetCreativeIds(creativeIds []int64) error {
    r.creativeIds = creativeIds
    r.Set("creative_ids", creativeIds)
    return nil
}

func (r TaobaoSimbaCreativesGetRequest) GetCreativeIds() []int64 {
    return r.creativeIds
}

func (r *TaobaoSimbaCreativesGetRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaCreativesGetRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

