package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelQuotaUpdateAPIRequest 库存更新接口 API请求
// taobao.xhotel.quota.update
//
// 库存更新接口
type TaobaoXhotelQuotaUpdateAPIRequest struct {
	model.Params
	// 修改日期列表
	_dates []string
	// 库存类型,0: 普通库存, 1: 普通保留房库存, 2:协议保留房库存
	_quotaType int64
	// room的gid
	_gid int64
	// 增减的值
	_quota int64
	// 数量类型, 2:增加房量,3:减少房量
	_quotaNumType int64
	// rate的id, rate库存时必填
	_rateId int64
	// 是否使用room库存,true使用，false不使用
	_useRoomInventory bool
}

// NewTaobaoXhotelQuotaUpdateRequest 初始化TaobaoXhotelQuotaUpdateAPIRequest对象
func NewTaobaoXhotelQuotaUpdateRequest() *TaobaoXhotelQuotaUpdateAPIRequest {
	return &TaobaoXhotelQuotaUpdateAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelQuotaUpdateAPIRequest) Reset() {
	r._dates = r._dates[:0]
	r._quotaType = 0
	r._gid = 0
	r._quota = 0
	r._quotaNumType = 0
	r._rateId = 0
	r._useRoomInventory = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.quota.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDates is Dates Setter
// 修改日期列表
func (r *TaobaoXhotelQuotaUpdateAPIRequest) SetDates(_dates []string) error {
	r._dates = _dates
	r.Set("dates", _dates)
	return nil
}

// GetDates Dates Getter
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetDates() []string {
	return r._dates
}

// SetQuotaType is QuotaType Setter
// 库存类型,0: 普通库存, 1: 普通保留房库存, 2:协议保留房库存
func (r *TaobaoXhotelQuotaUpdateAPIRequest) SetQuotaType(_quotaType int64) error {
	r._quotaType = _quotaType
	r.Set("quota_type", _quotaType)
	return nil
}

// GetQuotaType QuotaType Getter
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetQuotaType() int64 {
	return r._quotaType
}

// SetGid is Gid Setter
// room的gid
func (r *TaobaoXhotelQuotaUpdateAPIRequest) SetGid(_gid int64) error {
	r._gid = _gid
	r.Set("gid", _gid)
	return nil
}

// GetGid Gid Getter
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetGid() int64 {
	return r._gid
}

// SetQuota is Quota Setter
// 增减的值
func (r *TaobaoXhotelQuotaUpdateAPIRequest) SetQuota(_quota int64) error {
	r._quota = _quota
	r.Set("quota", _quota)
	return nil
}

// GetQuota Quota Getter
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetQuota() int64 {
	return r._quota
}

// SetQuotaNumType is QuotaNumType Setter
// 数量类型, 2:增加房量,3:减少房量
func (r *TaobaoXhotelQuotaUpdateAPIRequest) SetQuotaNumType(_quotaNumType int64) error {
	r._quotaNumType = _quotaNumType
	r.Set("quota_num_type", _quotaNumType)
	return nil
}

// GetQuotaNumType QuotaNumType Getter
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetQuotaNumType() int64 {
	return r._quotaNumType
}

// SetRateId is RateId Setter
// rate的id, rate库存时必填
func (r *TaobaoXhotelQuotaUpdateAPIRequest) SetRateId(_rateId int64) error {
	r._rateId = _rateId
	r.Set("rate_id", _rateId)
	return nil
}

// GetRateId RateId Getter
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetRateId() int64 {
	return r._rateId
}

// SetUseRoomInventory is UseRoomInventory Setter
// 是否使用room库存,true使用，false不使用
func (r *TaobaoXhotelQuotaUpdateAPIRequest) SetUseRoomInventory(_useRoomInventory bool) error {
	r._useRoomInventory = _useRoomInventory
	r.Set("use_room_inventory", _useRoomInventory)
	return nil
}

// GetUseRoomInventory UseRoomInventory Getter
func (r TaobaoXhotelQuotaUpdateAPIRequest) GetUseRoomInventory() bool {
	return r._useRoomInventory
}

var poolTaobaoXhotelQuotaUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelQuotaUpdateRequest()
	},
}

// GetTaobaoXhotelQuotaUpdateRequest 从 sync.Pool 获取 TaobaoXhotelQuotaUpdateAPIRequest
func GetTaobaoXhotelQuotaUpdateAPIRequest() *TaobaoXhotelQuotaUpdateAPIRequest {
	return poolTaobaoXhotelQuotaUpdateAPIRequest.Get().(*TaobaoXhotelQuotaUpdateAPIRequest)
}

// ReleaseTaobaoXhotelQuotaUpdateAPIRequest 将 TaobaoXhotelQuotaUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelQuotaUpdateAPIRequest(v *TaobaoXhotelQuotaUpdateAPIRequest) {
	v.Reset()
	poolTaobaoXhotelQuotaUpdateAPIRequest.Put(v)
}
