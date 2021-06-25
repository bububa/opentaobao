package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对推广组进行单独移动溢价 APIRequest
taobao.simba.adgroup.mobilediscount.update

对推广组进行单独移动溢价
*/
type TaobaoSimbaAdgroupMobilediscountUpdateRequest struct {
    model.Params

    // 推广组id数组(推广组id集合元素个数在1-200个之间，推广组id需要在同一个推广计划中)
    adgroupIds   []Number 

    // 折扣（折扣值在1-400之间）
    mobileDiscount   int64 

    // 昵称
    nick   string 

}

func NewTaobaoSimbaAdgroupMobilediscountUpdateRequest() *TaobaoSimbaAdgroupMobilediscountUpdateRequest{
    return &TaobaoSimbaAdgroupMobilediscountUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSimbaAdgroupMobilediscountUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.adgroup.mobilediscount.update"
}

func (r TaobaoSimbaAdgroupMobilediscountUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSimbaAdgroupMobilediscountUpdateRequest) SetAdgroupIds(adgroupIds []Number) error {
    r.adgroupIds = adgroupIds
    r.Set("adgroup_ids", adgroupIds)
    return nil
}

func (r TaobaoSimbaAdgroupMobilediscountUpdateRequest) GetAdgroupIds() []Number {
    return r.adgroupIds
}

func (r *TaobaoSimbaAdgroupMobilediscountUpdateRequest) SetMobileDiscount(mobileDiscount int64) error {
    r.mobileDiscount = mobileDiscount
    r.Set("mobile_discount", mobileDiscount)
    return nil
}

func (r TaobaoSimbaAdgroupMobilediscountUpdateRequest) GetMobileDiscount() int64 {
    return r.mobileDiscount
}

func (r *TaobaoSimbaAdgroupMobilediscountUpdateRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

func (r TaobaoSimbaAdgroupMobilediscountUpdateRequest) GetNick() string {
    return r.nick
}

