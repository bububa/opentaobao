package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest 查询单元离线列表30天转化周期 API请求
// taobao.subway.adgroup.offline.layeredfind
//
// 查询单元离线列表
type TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest struct {
	model.Params
	// 需要查询的单元id列表，不传表示不限制
	_adgroupIdIn []int64
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
	// 数据来源（pc站内：1；pc站外：2；无限站内：4；无限站内：5；销量明星：6）
	_pvTypeIn int64
	// 页码（0为第一页）
	_offset int64
	// 每页显示的记录条数
	_pageSize int64
	// 转化周期30-30天
	_effect int64
}

// NewTaobaoSubwayAdgroupOfflineLayeredfindRequest 初始化TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest对象
func NewTaobaoSubwayAdgroupOfflineLayeredfindRequest() *TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest {
	return &TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) Reset() {
	r._adgroupIdIn = r._adgroupIdIn[:0]
	r._startTime = ""
	r._endTime = ""
	r._pvTypeIn = 0
	r._offset = 0
	r._pageSize = 0
	r._effect = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.adgroup.offline.layeredfind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupIdIn is AdgroupIdIn Setter
// 需要查询的单元id列表，不传表示不限制
func (r *TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) SetAdgroupIdIn(_adgroupIdIn []int64) error {
	r._adgroupIdIn = _adgroupIdIn
	r.Set("adgroup_id_in", _adgroupIdIn)
	return nil
}

// GetAdgroupIdIn AdgroupIdIn Getter
func (r TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) GetAdgroupIdIn() []int64 {
	return r._adgroupIdIn
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetPvTypeIn is PvTypeIn Setter
// 数据来源（pc站内：1；pc站外：2；无限站内：4；无限站内：5；销量明星：6）
func (r *TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) SetPvTypeIn(_pvTypeIn int64) error {
	r._pvTypeIn = _pvTypeIn
	r.Set("pv_type_in", _pvTypeIn)
	return nil
}

// GetPvTypeIn PvTypeIn Getter
func (r TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) GetPvTypeIn() int64 {
	return r._pvTypeIn
}

// SetOffset is Offset Setter
// 页码（0为第一页）
func (r *TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 每页显示的记录条数
func (r *TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetEffect is Effect Setter
// 转化周期30-30天
func (r *TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) GetEffect() int64 {
	return r._effect
}

var poolTaobaoSubwayAdgroupOfflineLayeredfindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubwayAdgroupOfflineLayeredfindRequest()
	},
}

// GetTaobaoSubwayAdgroupOfflineLayeredfindRequest 从 sync.Pool 获取 TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest
func GetTaobaoSubwayAdgroupOfflineLayeredfindAPIRequest() *TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest {
	return poolTaobaoSubwayAdgroupOfflineLayeredfindAPIRequest.Get().(*TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest)
}

// ReleaseTaobaoSubwayAdgroupOfflineLayeredfindAPIRequest 将 TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubwayAdgroupOfflineLayeredfindAPIRequest(v *TaobaoSubwayAdgroupOfflineLayeredfindAPIRequest) {
	v.Reset()
	poolTaobaoSubwayAdgroupOfflineLayeredfindAPIRequest.Put(v)
}
