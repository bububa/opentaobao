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
type TaobaoXhotelQuotaUpdateAPIRequest struct {
    model.Params
    // 库存类型,0: 普通库存, 1: 普通保留房库存, 2:协议保留房库存
    _quotaType   int64
    // 是否使用room库存,true使用，false不使用
    _useRoomInventory   bool
    // room的gid
    _gid   int64
    // 增减的值
    _quota   int64
    // 数量类型, 2:增加房量,3:减少房量
    _quotaNumType   int64
    // 修改日期列表
    _dates   []string
    // rate的id, rate库存时必填
    _rateId   int64
}

// 初始化TaobaoXhotelQuotaUpdateAPIRequest对象
func NewTaobaoXhotelQuotaUpdateRequest() *TaobaoXhotelQuotaUpdateAPIRequest{
    return &TaobaoXhotelQuotaUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.quota.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QuotaType Setter
// 库存类型,0: 普通库存, 1: 普通保留房库存, 2:协议保留房库存
func (r *TaobaoXhotelQuotaUpdateAPIRequest) SetQuotaType(_quotaType int64) error {
    r._quotaType = _quotaType
    r.Set("quota_type", _quotaType)
    return nil
}

// QuotaType Getter
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetQuotaType() int64 {
    return r._quotaType
}
// UseRoomInventory Setter
// 是否使用room库存,true使用，false不使用
func (r *TaobaoXhotelQuotaUpdateAPIRequest) SetUseRoomInventory(_useRoomInventory bool) error {
    r._useRoomInventory = _useRoomInventory
    r.Set("use_room_inventory", _useRoomInventory)
    return nil
}

// UseRoomInventory Getter
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetUseRoomInventory() bool {
    return r._useRoomInventory
}
// Gid Setter
// room的gid
func (r *TaobaoXhotelQuotaUpdateAPIRequest) SetGid(_gid int64) error {
    r._gid = _gid
    r.Set("gid", _gid)
    return nil
}

// Gid Getter
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetGid() int64 {
    return r._gid
}
// Quota Setter
// 增减的值
func (r *TaobaoXhotelQuotaUpdateAPIRequest) SetQuota(_quota int64) error {
    r._quota = _quota
    r.Set("quota", _quota)
    return nil
}

// Quota Getter
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetQuota() int64 {
    return r._quota
}
// QuotaNumType Setter
// 数量类型, 2:增加房量,3:减少房量
func (r *TaobaoXhotelQuotaUpdateAPIRequest) SetQuotaNumType(_quotaNumType int64) error {
    r._quotaNumType = _quotaNumType
    r.Set("quota_num_type", _quotaNumType)
    return nil
}

// QuotaNumType Getter
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetQuotaNumType() int64 {
    return r._quotaNumType
}
// Dates Setter
// 修改日期列表
func (r *TaobaoXhotelQuotaUpdateAPIRequest) SetDates(_dates []string) error {
    r._dates = _dates
    r.Set("dates", _dates)
    return nil
}

// Dates Getter
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetDates() []string {
    return r._dates
}
// RateId Setter
// rate的id, rate库存时必填
func (r *TaobaoXhotelQuotaUpdateAPIRequest) SetRateId(_rateId int64) error {
    r._rateId = _rateId
    r.Set("rate_id", _rateId)
    return nil
}

// RateId Getter
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetRateId() int64 {
    return r._rateId
}
