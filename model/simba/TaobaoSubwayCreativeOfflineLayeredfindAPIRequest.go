package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCreativeOfflineLayeredfindAPIRequest 获取创意离线报表周期30天 API请求
// taobao.subway.creative.offline.layeredfind
//
// 获取创意离线报表
type TaobaoSubwayCreativeOfflineLayeredfindAPIRequest struct {
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
	// 转化周期30-30天
	_effect int64
}

// NewTaobaoSubwayCreativeOfflineLayeredfindRequest 初始化TaobaoSubwayCreativeOfflineLayeredfindAPIRequest对象
func NewTaobaoSubwayCreativeOfflineLayeredfindRequest() *TaobaoSubwayCreativeOfflineLayeredfindAPIRequest {
	return &TaobaoSubwayCreativeOfflineLayeredfindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.creative.offline.layeredfind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCreativeIdIn is CreativeIdIn Setter
// 需要查询的创意id列表，不传表示不限制
func (r *TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) SetCreativeIdIn(_creativeIdIn []int64) error {
	r._creativeIdIn = _creativeIdIn
	r.Set("creative_id_in", _creativeIdIn)
	return nil
}

// GetCreativeIdIn CreativeIdIn Getter
func (r TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) GetCreativeIdIn() []int64 {
	return r._creativeIdIn
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetPvTypeIn is PvTypeIn Setter
// 数据来源（pc站内：1；pc站外：2；无限站内：4；无限站内：5；销量明星：6）
func (r *TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) SetPvTypeIn(_pvTypeIn int64) error {
	r._pvTypeIn = _pvTypeIn
	r.Set("pv_type_in", _pvTypeIn)
	return nil
}

// GetPvTypeIn PvTypeIn Getter
func (r TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) GetPvTypeIn() int64 {
	return r._pvTypeIn
}

// SetCreativeIdEqual is CreativeIdEqual Setter
// 需要查询的创意id，不传表示不限
func (r *TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) SetCreativeIdEqual(_creativeIdEqual int64) error {
	r._creativeIdEqual = _creativeIdEqual
	r.Set("creative_id_equal", _creativeIdEqual)
	return nil
}

// GetCreativeIdEqual CreativeIdEqual Getter
func (r TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) GetCreativeIdEqual() int64 {
	return r._creativeIdEqual
}

// SetOffset is Offset Setter
// 页码（0为第一页）
func (r *TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 每页显示的记录条数
func (r *TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetEffect is Effect Setter
// 转化周期30-30天
func (r *TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaoSubwayCreativeOfflineLayeredfindAPIRequest) GetEffect() int64 {
	return r._effect
}
