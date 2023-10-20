package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwayadgroupofflinefindAPIRequest 查询单元离线多日汇总列表 API请求
// taobao.subway.adgroup.offline.find
//
// 查询单元离线列表
type TaobaosubwayadgroupofflinefindAPIRequest struct {
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

// NewTaobaosubwayadgroupofflinefindRequest 初始化TaobaosubwayadgroupofflinefindAPIRequest对象
func NewTaobaosubwayadgroupofflinefindRequest() *TaobaosubwayadgroupofflinefindAPIRequest {
	return &TaobaosubwayadgroupofflinefindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubwayadgroupofflinefindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.adgroup.offline.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubwayadgroupofflinefindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubwayadgroupofflinefindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupIdIn is AdgroupIdIn Setter
// 需要查询的单元id列表，不传表示不限制
func (r *TaobaosubwayadgroupofflinefindAPIRequest) SetAdgroupIdIn(_adgroupIdIn []int64) error {
	r._adgroupIdIn = _adgroupIdIn
	r.Set("adgroup_id_in", _adgroupIdIn)
	return nil
}

// GetAdgroupIdIn AdgroupIdIn Getter
func (r TaobaosubwayadgroupofflinefindAPIRequest) GetAdgroupIdIn() []int64 {
	return r._adgroupIdIn
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaosubwayadgroupofflinefindAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosubwayadgroupofflinefindAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaosubwayadgroupofflinefindAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosubwayadgroupofflinefindAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetPvTypeIn is PvTypeIn Setter
// 数据来源（pc站内：1；pc站外：2；无限站内：4；无限站内：5；销量明星：6）
func (r *TaobaosubwayadgroupofflinefindAPIRequest) SetPvTypeIn(_pvTypeIn int64) error {
	r._pvTypeIn = _pvTypeIn
	r.Set("pv_type_in", _pvTypeIn)
	return nil
}

// GetPvTypeIn PvTypeIn Getter
func (r TaobaosubwayadgroupofflinefindAPIRequest) GetPvTypeIn() int64 {
	return r._pvTypeIn
}

// SetOffset is Offset Setter
// 页码（0为第一页）
func (r *TaobaosubwayadgroupofflinefindAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaosubwayadgroupofflinefindAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 每页显示的记录条数
func (r *TaobaosubwayadgroupofflinefindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosubwayadgroupofflinefindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetEffect is Effect Setter
// 转化周期-1-15累计天数，1-1转化天数，3-3转化天数，7-7转化天数，15-15转化天数，不传默认为15累计天数
func (r *TaobaosubwayadgroupofflinefindAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaosubwayadgroupofflinefindAPIRequest) GetEffect() int64 {
	return r._effect
}

// SetCampaignIdEqual is CampaignIdEqual Setter
// 需要查询的计划id，不传表示不限制
func (r *TaobaosubwayadgroupofflinefindAPIRequest) SetCampaignIdEqual(_campaignIdEqual int64) error {
	r._campaignIdEqual = _campaignIdEqual
	r.Set("campaign_id_equal", _campaignIdEqual)
	return nil
}

// GetCampaignIdEqual CampaignIdEqual Getter
func (r TaobaosubwayadgroupofflinefindAPIRequest) GetCampaignIdEqual() int64 {
	return r._campaignIdEqual
}
