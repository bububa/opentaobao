package simba

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销量明星更新一个推广组的信息 API请求
taobao.simba.salestar.adgroup.update

更新一个推广组的信息，可以设置 是否上线
*/
type TaobaoSimbaSalestarAdgroupUpdateRequest struct {
    model.Params
    // 用户设置的上下线状态 offline-下线(暂停竞价)； online-上线；默认为online
    onlineStatus   string
    // 推广组Id
    adgroupId   int64
}

// 初始化TaobaoSimbaSalestarAdgroupUpdateRequest对象
func NewTaobaoSimbaSalestarAdgroupUpdateRequest() *TaobaoSimbaSalestarAdgroupUpdateRequest{
    return &TaobaoSimbaSalestarAdgroupUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarAdgroupUpdateRequest) GetApiMethodName() string {
    return "taobao.simba.salestar.adgroup.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarAdgroupUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OnlineStatus Setter
// 用户设置的上下线状态 offline-下线(暂停竞价)； online-上线；默认为online
func (r *TaobaoSimbaSalestarAdgroupUpdateRequest) SetOnlineStatus(onlineStatus string) error {
    r.onlineStatus = onlineStatus
    r.Set("online_status", onlineStatus)
    return nil
}

// OnlineStatus Getter
func (r TaobaoSimbaSalestarAdgroupUpdateRequest) GetOnlineStatus() string {
    return r.onlineStatus
}
// AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaSalestarAdgroupUpdateRequest) SetAdgroupId(adgroupId int64) error {
    r.adgroupId = adgroupId
    r.Set("adgroup_id", adgroupId)
    return nil
}

// AdgroupId Getter
func (r TaobaoSimbaSalestarAdgroupUpdateRequest) GetAdgroupId() int64 {
    return r.adgroupId
}
