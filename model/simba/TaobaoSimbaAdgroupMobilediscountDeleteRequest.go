package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量删除adgroup的移动溢价 API请求
taobao.simba.adgroup.mobilediscount.delete

批量删除adgroup的移动溢价
*/
type TaobaoSimbaAdgroupMobilediscountDeleteRequest struct {
    model.Params
    // 昵称
    _nick   string
    // adgroup主键数组（批量最多支持200个）
    _adgroupIds   []int64
}

// 初始化TaobaoSimbaAdgroupMobilediscountDeleteRequest对象
func NewTaobaoSimbaAdgroupMobilediscountDeleteRequest() *TaobaoSimbaAdgroupMobilediscountDeleteRequest{
    return &TaobaoSimbaAdgroupMobilediscountDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupMobilediscountDeleteRequest) GetApiMethodName() string {
    return "taobao.simba.adgroup.mobilediscount.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupMobilediscountDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 昵称
func (r *TaobaoSimbaAdgroupMobilediscountDeleteRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaAdgroupMobilediscountDeleteRequest) GetNick() string {
    return r._nick
}
// AdgroupIds Setter
// adgroup主键数组（批量最多支持200个）
func (r *TaobaoSimbaAdgroupMobilediscountDeleteRequest) SetAdgroupIds(_adgroupIds []int64) error {
    r._adgroupIds = _adgroupIds
    r.Set("adgroup_ids", _adgroupIds)
    return nil
}

// AdgroupIds Getter
func (r TaobaoSimbaAdgroupMobilediscountDeleteRequest) GetAdgroupIds() []int64 {
    return r._adgroupIds
}
