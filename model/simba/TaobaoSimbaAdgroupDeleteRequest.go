package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
删除一个推广组 API请求
taobao.simba.adgroup.delete

删除一个推广组
*/
type TaobaoSimbaAdgroupDeleteRequest struct {
    model.Params
    // 主人昵称
    _nick   string
    // 推广组Id
    _adgroupId   int64
}

// 初始化TaobaoSimbaAdgroupDeleteRequest对象
func NewTaobaoSimbaAdgroupDeleteRequest() *TaobaoSimbaAdgroupDeleteRequest{
    return &TaobaoSimbaAdgroupDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupDeleteRequest) GetApiMethodName() string {
    return "taobao.simba.adgroup.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupDeleteRequest) SetNick(_nick string) error {
    r._nick = _nick
    r.Set("nick", _nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaAdgroupDeleteRequest) GetNick() string {
    return r._nick
}
// AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaAdgroupDeleteRequest) SetAdgroupId(_adgroupId int64) error {
    r._adgroupId = _adgroupId
    r.Set("adgroup_id", _adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaAdgroupDeleteRequest) GetAdgroupId() int64 {
    return r._adgroupId
}
