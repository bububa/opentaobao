package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCreativeofflineFindAPIRequest 获取创意离线多日汇总报表 API请求
// taobao.subway.creativeoffline.find
//
// 获取创意离线报表
type TaobaoSubwayCreativeofflineFindAPIRequest struct {
	model.Params
	// 需要查询的创意id列表，不传表示不限制
	_creativeIdIn []int64
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
	// 数据来源（pc站内：1；pc站外：2；无限站内：4；无限站内：5；销量明星：6）
	_pvTypeIn int64
	// 需要查询的创意id，不传表示不限
	_creativeIdEqual int64
	// 页码（0为第一页）
	_offset int64
	// 每页显示的记录条数
	_pageSize int64
	// 转化周期-1-15累计天数，1-1转化天数，3-3转化天数，7-7转化天数，15-15转化天数，不传默认为15累计天数
	_effect int64
	// 需要查询的计划id，不传表示不限制
	_campaignIdEqual int64
}

// NewTaobaoSubwayCreativeofflineFindRequest 初始化TaobaoSubwayCreativeofflineFindAPIRequest对象
func NewTaobaoSubwayCreativeofflineFindRequest() *TaobaoSubwayCreativeofflineFindAPIRequest {
	return &TaobaoSubwayCreativeofflineFindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayCreativeofflineFindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.creativeoffline.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayCreativeofflineFindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayCreativeofflineFindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreativeIdIn is CreativeIdIn Setter
// 需要查询的创意id列表，不传表示不限制
func (r *TaobaoSubwayCreativeofflineFindAPIRequest) SetCreativeIdIn(_creativeIdIn []int64) error {
	r._creativeIdIn = _creativeIdIn
	r.Set("creative_id_in", _creativeIdIn)
	return nil
}

// GetCreativeIdIn CreativeIdIn Getter
func (r TaobaoSubwayCreativeofflineFindAPIRequest) GetCreativeIdIn() []int64 {
	return r._creativeIdIn
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaoSubwayCreativeofflineFindAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSubwayCreativeofflineFindAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoSubwayCreativeofflineFindAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoSubwayCreativeofflineFindAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetPvTypeIn is PvTypeIn Setter
// 数据来源（pc站内：1；pc站外：2；无限站内：4；无限站内：5；销量明星：6）
func (r *TaobaoSubwayCreativeofflineFindAPIRequest) SetPvTypeIn(_pvTypeIn int64) error {
	r._pvTypeIn = _pvTypeIn
	r.Set("pv_type_in", _pvTypeIn)
	return nil
}

// GetPvTypeIn PvTypeIn Getter
func (r TaobaoSubwayCreativeofflineFindAPIRequest) GetPvTypeIn() int64 {
	return r._pvTypeIn
}

// SetCreativeIdEqual is CreativeIdEqual Setter
// 需要查询的创意id，不传表示不限
func (r *TaobaoSubwayCreativeofflineFindAPIRequest) SetCreativeIdEqual(_creativeIdEqual int64) error {
	r._creativeIdEqual = _creativeIdEqual
	r.Set("creative_id_equal", _creativeIdEqual)
	return nil
}

// GetCreativeIdEqual CreativeIdEqual Getter
func (r TaobaoSubwayCreativeofflineFindAPIRequest) GetCreativeIdEqual() int64 {
	return r._creativeIdEqual
}

// SetOffset is Offset Setter
// 页码（0为第一页）
func (r *TaobaoSubwayCreativeofflineFindAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaoSubwayCreativeofflineFindAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 每页显示的记录条数
func (r *TaobaoSubwayCreativeofflineFindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSubwayCreativeofflineFindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetEffect is Effect Setter
// 转化周期-1-15累计天数，1-1转化天数，3-3转化天数，7-7转化天数，15-15转化天数，不传默认为15累计天数
func (r *TaobaoSubwayCreativeofflineFindAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaoSubwayCreativeofflineFindAPIRequest) GetEffect() int64 {
	return r._effect
}

// SetCampaignIdEqual is CampaignIdEqual Setter
// 需要查询的计划id，不传表示不限制
func (r *TaobaoSubwayCreativeofflineFindAPIRequest) SetCampaignIdEqual(_campaignIdEqual int64) error {
	r._campaignIdEqual = _campaignIdEqual
	r.Set("campaign_id_equal", _campaignIdEqual)
	return nil
}

// GetCampaignIdEqual CampaignIdEqual Getter
func (r TaobaoSubwayCreativeofflineFindAPIRequest) GetCampaignIdEqual() int64 {
	return r._campaignIdEqual
}
