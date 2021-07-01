package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
对推广组进行单独移动溢价 API请求
taobao.simba.adgroup.mobilediscount.update

对推广组进行单独移动溢价
*/
type TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest struct {
    model.Params
    // 推广组id数组(推广组id集合元素个数在1-200个之间，推广组id需要在同一个推广计划中)
    _adgroupIds   []int64
    // 折扣（折扣值在1-400之间）
    _mobileDiscount   int64
    // 昵称
    _nick   string
}

// 初始化TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest对象
func NewTaobaoSimbaAdgroupMobilediscountUpdateRequest() *TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest{
    return &TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.simba.adgroup.mobilediscount.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AdgroupIds Setter
// 推广组id数组(推广组id集合元素个数在1-200个之间，推广组id需要在同一个推广计划中)
func (r *TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) SetAdgroupIds(_adgroupIds []int64) error {
    r._adgroupIds = _adgroupIds
    r.Set("adgroup_ids", _adgroupIds)
    return nil
}

// AdgroupIds Getter
func (r TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) GetAdgroupIds() []int64 {
    return r._adgroupIds
}
// MobileDiscount Setter
// 折扣（折扣值在1-400之间）
func (r *TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) SetMobileDiscount(_mobileDiscount int64) error {
    r._mobileDiscount = _mobileDiscount
    r.Set("mobile_discount", _mobileDiscount)
    return nil
}

// MobileDiscount Getter
func (r TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) GetMobileDiscount() int64 {
    return r._mobileDiscount
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) GetNick() string {
    return r._nick
}
