package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwaycrowdofflinelayeredfindAPIRequest 获取人群离线报表30转化周期 API请求
// taobao.subway.crowd.offline.layeredfind
//
// 获取人群离线报表
type TaobaosubwaycrowdofflinelayeredfindAPIRequest struct {
	model.Params
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
	// 数据来源（pc站内：1；pc站外：2；无限站内：4；无限站内：5；销量明星：6）
	_pvTypeIn int64
	// 需要查询的创意id，不传表示不限
	_crowdIdEqual int64
	// 页码（0为第一页）
	_offset int64
	// 每页显示的记录条数
	_pageSize int64
	// 转化周期30-30天
	_effect int64
}

// NewTaobaosubwaycrowdofflinelayeredfindRequest 初始化TaobaosubwaycrowdofflinelayeredfindAPIRequest对象
func NewTaobaosubwaycrowdofflinelayeredfindRequest() *TaobaosubwaycrowdofflinelayeredfindAPIRequest {
	return &TaobaosubwaycrowdofflinelayeredfindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubwaycrowdofflinelayeredfindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.crowd.offline.layeredfind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubwaycrowdofflinelayeredfindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubwaycrowdofflinelayeredfindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaosubwaycrowdofflinelayeredfindAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaosubwaycrowdofflinelayeredfindAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaosubwaycrowdofflinelayeredfindAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaosubwaycrowdofflinelayeredfindAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetPvTypeIn is PvTypeIn Setter
// 数据来源（pc站内：1；pc站外：2；无限站内：4；无限站内：5；销量明星：6）
func (r *TaobaosubwaycrowdofflinelayeredfindAPIRequest) SetPvTypeIn(_pvTypeIn int64) error {
	r._pvTypeIn = _pvTypeIn
	r.Set("pv_type_in", _pvTypeIn)
	return nil
}

// GetPvTypeIn PvTypeIn Getter
func (r TaobaosubwaycrowdofflinelayeredfindAPIRequest) GetPvTypeIn() int64 {
	return r._pvTypeIn
}

// SetCrowdIdEqual is CrowdIdEqual Setter
// 需要查询的创意id，不传表示不限
func (r *TaobaosubwaycrowdofflinelayeredfindAPIRequest) SetCrowdIdEqual(_crowdIdEqual int64) error {
	r._crowdIdEqual = _crowdIdEqual
	r.Set("crowd_id_equal", _crowdIdEqual)
	return nil
}

// GetCrowdIdEqual CrowdIdEqual Getter
func (r TaobaosubwaycrowdofflinelayeredfindAPIRequest) GetCrowdIdEqual() int64 {
	return r._crowdIdEqual
}

// SetOffset is Offset Setter
// 页码（0为第一页）
func (r *TaobaosubwaycrowdofflinelayeredfindAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaosubwaycrowdofflinelayeredfindAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 每页显示的记录条数
func (r *TaobaosubwaycrowdofflinelayeredfindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaosubwaycrowdofflinelayeredfindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetEffect is Effect Setter
// 转化周期30-30天
func (r *TaobaosubwaycrowdofflinelayeredfindAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaosubwaycrowdofflinelayeredfindAPIRequest) GetEffect() int64 {
	return r._effect
}
