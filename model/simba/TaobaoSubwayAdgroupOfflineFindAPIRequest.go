package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayAdgroupOfflineFindAPIRequest 查询单元离线多日汇总列表 API请求
// taobao.subway.adgroup.offline.find
//
// 查询单元离线列表
type TaobaoSubwayAdgroupOfflineFindAPIRequest struct {
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
	// 转化周期-1-15累计天数，1-1转化天数，3-3转化天数，7-7转化天数，15-15转化天数，不传默认为15累计天数
	_effect int64
	// 需要查询的计划id，不传表示不限制
	_campaignIdEqual int64
}

// NewTaobaoSubwayAdgroupOfflineFindRequest 初始化TaobaoSubwayAdgroupOfflineFindAPIRequest对象
func NewTaobaoSubwayAdgroupOfflineFindRequest() *TaobaoSubwayAdgroupOfflineFindAPIRequest {
	return &TaobaoSubwayAdgroupOfflineFindAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubwayAdgroupOfflineFindAPIRequest) Reset() {
	r._adgroupIdIn = r._adgroupIdIn[:0]
	r._startTime = ""
	r._endTime = ""
	r._pvTypeIn = 0
	r._offset = 0
	r._pageSize = 0
	r._effect = 0
	r._campaignIdEqual = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayAdgroupOfflineFindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.adgroup.offline.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayAdgroupOfflineFindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayAdgroupOfflineFindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupIdIn is AdgroupIdIn Setter
// 需要查询的单元id列表，不传表示不限制
func (r *TaobaoSubwayAdgroupOfflineFindAPIRequest) SetAdgroupIdIn(_adgroupIdIn []int64) error {
	r._adgroupIdIn = _adgroupIdIn
	r.Set("adgroup_id_in", _adgroupIdIn)
	return nil
}

// GetAdgroupIdIn AdgroupIdIn Getter
func (r TaobaoSubwayAdgroupOfflineFindAPIRequest) GetAdgroupIdIn() []int64 {
	return r._adgroupIdIn
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaoSubwayAdgroupOfflineFindAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSubwayAdgroupOfflineFindAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoSubwayAdgroupOfflineFindAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoSubwayAdgroupOfflineFindAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetPvTypeIn is PvTypeIn Setter
// 数据来源（pc站内：1；pc站外：2；无限站内：4；无限站内：5；销量明星：6）
func (r *TaobaoSubwayAdgroupOfflineFindAPIRequest) SetPvTypeIn(_pvTypeIn int64) error {
	r._pvTypeIn = _pvTypeIn
	r.Set("pv_type_in", _pvTypeIn)
	return nil
}

// GetPvTypeIn PvTypeIn Getter
func (r TaobaoSubwayAdgroupOfflineFindAPIRequest) GetPvTypeIn() int64 {
	return r._pvTypeIn
}

// SetOffset is Offset Setter
// 页码（0为第一页）
func (r *TaobaoSubwayAdgroupOfflineFindAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaoSubwayAdgroupOfflineFindAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 每页显示的记录条数
func (r *TaobaoSubwayAdgroupOfflineFindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSubwayAdgroupOfflineFindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetEffect is Effect Setter
// 转化周期-1-15累计天数，1-1转化天数，3-3转化天数，7-7转化天数，15-15转化天数，不传默认为15累计天数
func (r *TaobaoSubwayAdgroupOfflineFindAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaoSubwayAdgroupOfflineFindAPIRequest) GetEffect() int64 {
	return r._effect
}

// SetCampaignIdEqual is CampaignIdEqual Setter
// 需要查询的计划id，不传表示不限制
func (r *TaobaoSubwayAdgroupOfflineFindAPIRequest) SetCampaignIdEqual(_campaignIdEqual int64) error {
	r._campaignIdEqual = _campaignIdEqual
	r.Set("campaign_id_equal", _campaignIdEqual)
	return nil
}

// GetCampaignIdEqual CampaignIdEqual Getter
func (r TaobaoSubwayAdgroupOfflineFindAPIRequest) GetCampaignIdEqual() int64 {
	return r._campaignIdEqual
}

var poolTaobaoSubwayAdgroupOfflineFindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubwayAdgroupOfflineFindRequest()
	},
}

// GetTaobaoSubwayAdgroupOfflineFindRequest 从 sync.Pool 获取 TaobaoSubwayAdgroupOfflineFindAPIRequest
func GetTaobaoSubwayAdgroupOfflineFindAPIRequest() *TaobaoSubwayAdgroupOfflineFindAPIRequest {
	return poolTaobaoSubwayAdgroupOfflineFindAPIRequest.Get().(*TaobaoSubwayAdgroupOfflineFindAPIRequest)
}

// ReleaseTaobaoSubwayAdgroupOfflineFindAPIRequest 将 TaobaoSubwayAdgroupOfflineFindAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubwayAdgroupOfflineFindAPIRequest(v *TaobaoSubwayAdgroupOfflineFindAPIRequest) {
	v.Reset()
	poolTaobaoSubwayAdgroupOfflineFindAPIRequest.Put(v)
}
