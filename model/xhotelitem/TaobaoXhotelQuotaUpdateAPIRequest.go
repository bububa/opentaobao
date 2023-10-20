package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelquotaupdateAPIRequest 库存更新接口 API请求
// taobao.xhotel.quota.update
//
// 库存更新接口
type TaobaoxhotelquotaupdateAPIRequest struct {
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

// NewTaobaoxhotelquotaupdateRequest 初始化TaobaoxhotelquotaupdateAPIRequest对象
func NewTaobaoxhotelquotaupdateRequest() *TaobaoxhotelquotaupdateAPIRequest {
	return &TaobaoxhotelquotaupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelquotaupdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.quota.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelquotaupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelquotaupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDates is Dates Setter
// 修改日期列表
func (r *TaobaoxhotelquotaupdateAPIRequest) SetDates(_dates []string) error {
	r._dates = _dates
	r.Set("dates", _dates)
	return nil
}

// GetDates Dates Getter
func (r TaobaoxhotelquotaupdateAPIRequest) GetDates() []string {
	return r._dates
}

// SetQuotaType is QuotaType Setter
// 库存类型,0: 普通库存, 1: 普通保留房库存, 2:协议保留房库存
func (r *TaobaoxhotelquotaupdateAPIRequest) SetQuotaType(_quotaType int64) error {
	r._quotaType = _quotaType
	r.Set("quota_type", _quotaType)
	return nil
}

// GetQuotaType QuotaType Getter
func (r TaobaoxhotelquotaupdateAPIRequest) GetQuotaType() int64 {
	return r._quotaType
}

// SetGid is Gid Setter
// room的gid
func (r *TaobaoxhotelquotaupdateAPIRequest) SetGid(_gid int64) error {
	r._gid = _gid
	r.Set("gid", _gid)
	return nil
}

// GetGid Gid Getter
func (r TaobaoxhotelquotaupdateAPIRequest) GetGid() int64 {
	return r._gid
}

// SetQuota is Quota Setter
// 增减的值
func (r *TaobaoxhotelquotaupdateAPIRequest) SetQuota(_quota int64) error {
	r._quota = _quota
	r.Set("quota", _quota)
	return nil
}

// GetQuota Quota Getter
func (r TaobaoxhotelquotaupdateAPIRequest) GetQuota() int64 {
	return r._quota
}

// SetQuotaNumType is QuotaNumType Setter
// 数量类型, 2:增加房量,3:减少房量
func (r *TaobaoxhotelquotaupdateAPIRequest) SetQuotaNumType(_quotaNumType int64) error {
	r._quotaNumType = _quotaNumType
	r.Set("quota_num_type", _quotaNumType)
	return nil
}

// GetQuotaNumType QuotaNumType Getter
func (r TaobaoxhotelquotaupdateAPIRequest) GetQuotaNumType() int64 {
	return r._quotaNumType
}

// SetRateId is RateId Setter
// rate的id, rate库存时必填
func (r *TaobaoxhotelquotaupdateAPIRequest) SetRateId(_rateId int64) error {
	r._rateId = _rateId
	r.Set("rate_id", _rateId)
	return nil
}

// GetRateId RateId Getter
func (r TaobaoxhotelquotaupdateAPIRequest) GetRateId() int64 {
	return r._rateId
}

// SetUseRoomInventory is UseRoomInventory Setter
// 是否使用room库存,true使用，false不使用
func (r *TaobaoxhotelquotaupdateAPIRequest) SetUseRoomInventory(_useRoomInventory bool) error {
	r._useRoomInventory = _useRoomInventory
	r.Set("use_room_inventory", _useRoomInventory)
	return nil
}

// GetUseRoomInventory UseRoomInventory Getter
func (r TaobaoxhotelquotaupdateAPIRequest) GetUseRoomInventory() bool {
	return r._useRoomInventory
}
