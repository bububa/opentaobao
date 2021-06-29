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
    nick   string
    // 推广组Id
    adgroupId   int64
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
func (r *TaobaoSimbaAdgroupDeleteRequest) SetNick(nick string) error {
    r.nick = nick
    r.Set("nick", nick)
    return nil
}

// Nick Getter
func (r TaobaoSimbaAdgroupDeleteRequest) GetNick() string {
    return r.nick
}
// AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaAdgroupDeleteRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaAdgroupDeleteRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
