package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存更新接口 API请求
taobao.xhotel.quota.update

库存更新接口
*/
type TaobaoXhotelQuotaUpdateRequest struct {
    model.Params
    // 库存类型,0: 普通库存, 1: 普通保留房库存, 2:协议保留房库存
    quotaType   int64
    // 是否使用room库存,true使用，false不使用
    useRoomInventory   bool
    // room的gid
    gid   int64
    // 增减的值
    quota   int64
    // 数量类型, 2:增加房量,3:减少房量
    quotaNumType   int64
    // 修改日期列表
    dates   []string
    // rate的id, rate库存时必填
    rateId   int64
}

// 初始化TaobaoXhotelQuotaUpdateRequest对象
func NewTaobaoXhotelQuotaUpdateRequest() *TaobaoXhotelQuotaUpdateRequest{
    return &TaobaoXhotelQuotaUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelQuotaUpdateRequest) GetApiMethodName() string {
    return "taobao.xhotel.quota.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelQuotaUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QuotaType Setter
// 库存类型,0: 普通库存, 1: 普通保留房库存, 2:协议保留房库存
func (r *TaobaoXhotelQuotaUpdateRequest) SetQuotaType(quotaType int64) error {
    r.quotaType = quotaType
    r.Set("quota_type", quotaType)
    return nil
}

// QuotaType Getter
func (r TaobaoXhotelQuotaUpdateRequest) GetQuotaType() int64 {
    return r.quotaType
}
// UseRoomInventory Setter
// 是否使用room库存,true使用，false不使用
func (r *TaobaoXhotelQuotaUpdateRequest) SetUseRoomInventory(useRoomInventory bool) error {
    r.useRoomInventory = useRoomInventory
    r.Set("use_room_inventory", useRoomInventory)
    return nil
}

// UseRoomInventory Getter
func (r TaobaoXhotelQuotaUpdateRequest) GetUseRoomInventory() bool {
    return r.useRoomInventory
}
// Gid Setter
// room的gid
func (r *TaobaoXhotelQuotaUpdateRequest) SetGid(gid int64) error {
    r.gid = gid
    r.Set("gid", gid)
    return nil
}

// Gid Getter
func (r TaobaoXhotelQuotaUpdateRequest) GetGid() int64 {
    return r.gid
}
// Quota Setter
// 增减的值
func (r *TaobaoXhotelQuotaUpdateRequest) SetQuota(quota int64) error {
    r.quota = quota
    r.Set("quota", quota)
    return nil
}

// Quota Getter
func (r TaobaoXhotelQuotaUpdateRequest) GetQuota() int64 {
    return r.quota
}
// QuotaNumType Setter
// 数量类型, 2:增加房量,3:减少房量
func (r *TaobaoXhotelQuotaUpdateRequest) SetQuotaNumType(quotaNumType int64) error {
    r.quotaNumType = quotaNumType
    r.Set("quota_num_type", quotaNumType)
    return nil
}

// QuotaNumType Getter
func (r TaobaoXhotelQuotaUpdateRequest) GetQuotaNumType() int64 {
    return r.quotaNumType
}
// Dates Setter
// 修改日期列表
func (r *TaobaoXhotelQuotaUpdateRequest) SetDates(dates []string) error {
    r.dates = dates
    r.Set("dates", dates)
    return nil
}

// Dates Getter
func (r TaobaoXhotelQuotaUpdateRequest) GetDates() []string {
    return r.dates
}
// RateId Setter
// rate的id, rate库存时必填
func (r *TaobaoXhotelQuotaUpdateRequest) SetRateId(rateId int64) error {
    r.rateId = rateId
    r.Set("rate_id", rateId)
    return nil
}

// RateId Getter
func (r TaobaoXhotelQuotaUpdateRequest) GetRateId() int64 {
    return r.rateId
}
