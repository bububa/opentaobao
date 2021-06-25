package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广组的信息 APIRequest
taobao.simba.adgroup.update

更新一个推广组的信息，可以设置默认出价、是否上线、非搜索出价、非搜索是否使用默认出价
*/
type TaobaoSimbaAdgroupUpdateRequest struct {
    model.Params

    // 非搜索是否使用默认出价，false-不用；true-使用；默认为true;
    useNonsearchDefaultPrice   string 

    // 用户设置的上下线状态 offline-下线(暂停竞价)； online-上线；默认为online
    onlineStatus   string 

    // 推广组Id
    adgroupId   int64 

    // 默认出价，单位是分，不能小于5
    defaultPrice   int64 

    // 非搜索出价，单位是分，不能小于5，如果use_nonseatch_default_price为使用默认出价，则此nonsearch_max_price字段传入的数据不起作用，商品将使用默认非搜索出价
    nonsearchMaxPrice   int64 

    // 主人昵称
    nick   string 

}

func NewTaobaoSimbaAdgroupUpdateRequest() *TaobaoSimbaAdgroupUpdateRequest{
    return &TaobaoSimbaAdgroupUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaAdgroupUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.adgroup.update"
}

func (r TaobaoSimbaAdgroupUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaAdgroupUpdateRequest) SetUseNonsearchDefaultPrice(useNonsearchDefaultPrice string) error {
    r.useNonsearchDefaultPrice = useNonsearchDefaultPrice
    r.Set("use_nonsearch_default_price", useNonsearchDefaultPrice)
    return nil
}

func (r TaobaoSimbaAdgroupUpdateRequest) GetUseNonsearchDefaultPrice() string {
    return r.useNonsearchDefaultPrice
}

func (r *TaobaoSimbaAdgroupUpdateRequest) SetOnlineStatus(onlineStatus string) error {
    r.onlineStatus = onlineStatus
    r.Set("online_status", onlineStatus)
    return nil
}

func (r TaobaoSimbaAdgroupUpdateRequest) GetOnlineStatus() string {
    return r.onlineStatus
}

func (r *TaobaoSimbaAdgroupUpdateRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

func (r TaobaoSimbaAdgroupUpdateRequest) GetAdgroupId() int64 {
    return r.adgroupId
}

func (r *TaobaoSimbaAdgroupUpdateRequest) SetDefaultPrice(defaultPrice int64) error {
    r.defaultPrice = defaultPrice
    r.Set("default_price", defaultPrice)
    return nil
}

func (r TaobaoSimbaAdgroupUpdateRequest) GetDefaultPrice() int64 {
    return r.defaultPrice
}

func (r *TaobaoSimbaAdgroupUpdateRequest) SetNonsearchMaxPrice(nonsearchMaxPrice int64) error {
    r.nonsearchMaxPrice = nonsearchMaxPrice
    r.Set("nonsearch_max_price", nonsearchMaxPrice)
    return nil
}

func (r TaobaoSimbaAdgroupUpdateRequest) GetNonsearchMaxPrice() int64 {
    return r.nonsearchMaxPrice
}

func (r *TaobaoSimbaAdgroupUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaAdgroupUpdateRequest) GetNick() string {
    return r.nick
}

