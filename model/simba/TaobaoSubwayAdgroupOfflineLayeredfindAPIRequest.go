package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwayadgroupofflinelayeredfindAPIRequest 查询单元离线列表30天转化周期 API请求
// taobao.subway.adgroup.offline.layeredfind
//
// 查询单元离线列表
type TaobaosubwayadgroupofflinelayeredfindAPIRequest struct {
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

// NewTaobaosubwayadgroupofflinelayeredfindRequest 初始化TaobaosubwayadgroupofflinelayeredfindAPIRequest对象
func NewTaobaosubwayadgroupofflinelayeredfindRequest() *TaobaosubwayadgroupofflinelayeredfindAPIRequest {
	return &TaobaosubwayadgroupofflinelayeredfindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubwayadgroupofflinelayeredfindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.adgroup.offline.layeredfind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubwayadgroupofflinelayeredfindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubwayadgroupofflinelayeredfindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupIdIn is AdgroupIdIn Setter
// 需要查询的单元id列表，不传表示不限制
func (r *TaobaosubwayadgroupofflinelayeredfindAPIRequest) SetAdgroupIdIn(_adgroupIdIn []int64) error {
	r._adgroupIdIn = _adgroupIdIn
	r.Set("adgroup_id_in", _adgroupIdIn)
	return nil
}

// GetAdgroupIdIn AdgroupIdIn Getter
func (r TaobaosubwayadgroupofflinelayeredfindAPIRequest) GetAdgroupIdIn() []int64 {
	return r._adgroupIdIn
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaosubwayadgroupofflinelayeredfindAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosubwayadgroupofflinelayeredfindAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaosubwayadgroupofflinelayeredfindAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosubwayadgroupofflinelayeredfindAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetPvTypeIn is PvTypeIn Setter
// 数据来源（pc站内：1；pc站外：2；无限站内：4；无限站内：5；销量明星：6）
func (r *TaobaosubwayadgroupofflinelayeredfindAPIRequest) SetPvTypeIn(_pvTypeIn int64) error {
	r._pvTypeIn = _pvTypeIn
	r.Set("pv_type_in", _pvTypeIn)
	return nil
}

// GetPvTypeIn PvTypeIn Getter
func (r TaobaosubwayadgroupofflinelayeredfindAPIRequest) GetPvTypeIn() int64 {
	return r._pvTypeIn
}

// SetOffset is Offset Setter
// 页码（0为第一页）
func (r *TaobaosubwayadgroupofflinelayeredfindAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaosubwayadgroupofflinelayeredfindAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 每页显示的记录条数
func (r *TaobaosubwayadgroupofflinelayeredfindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosubwayadgroupofflinelayeredfindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetEffect is Effect Setter
// 转化周期30-30天
func (r *TaobaosubwayadgroupofflinelayeredfindAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaosubwayadgroupofflinelayeredfindAPIRequest) GetEffect() int64 {
	return r._effect
}
