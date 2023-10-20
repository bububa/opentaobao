package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCrowdOfflineLayeredfindAPIRequest 获取人群离线报表30转化周期 API请求
// taobao.subway.crowd.offline.layeredfind
//
// 获取人群离线报表
type TaobaoSubwayCrowdOfflineLayeredfindAPIRequest struct {
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

// NewTaobaoSubwayCrowdOfflineLayeredfindRequest 初始化TaobaoSubwayCrowdOfflineLayeredfindAPIRequest对象
func NewTaobaoSubwayCrowdOfflineLayeredfindRequest() *TaobaoSubwayCrowdOfflineLayeredfindAPIRequest {
	return &TaobaoSubwayCrowdOfflineLayeredfindAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) Reset() {
	r._startTime = ""
	r._endTime = ""
	r._pvTypeIn = 0
	r._crowdIdEqual = 0
	r._offset = 0
	r._pageSize = 0
	r._effect = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.crowd.offline.layeredfind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetPvTypeIn is PvTypeIn Setter
// 数据来源（pc站内：1；pc站外：2；无限站内：4；无限站内：5；销量明星：6）
func (r *TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) SetPvTypeIn(_pvTypeIn int64) error {
	r._pvTypeIn = _pvTypeIn
	r.Set("pv_type_in", _pvTypeIn)
	return nil
}

// GetPvTypeIn PvTypeIn Getter
func (r TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) GetPvTypeIn() int64 {
	return r._pvTypeIn
}

// SetCrowdIdEqual is CrowdIdEqual Setter
// 需要查询的创意id，不传表示不限
func (r *TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) SetCrowdIdEqual(_crowdIdEqual int64) error {
	r._crowdIdEqual = _crowdIdEqual
	r.Set("crowd_id_equal", _crowdIdEqual)
	return nil
}

// GetCrowdIdEqual CrowdIdEqual Getter
func (r TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) GetCrowdIdEqual() int64 {
	return r._crowdIdEqual
}

// SetOffset is Offset Setter
// 页码（0为第一页）
func (r *TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 每页显示的记录条数
func (r *TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetEffect is Effect Setter
// 转化周期30-30天
func (r *TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) GetEffect() int64 {
	return r._effect
}

var poolTaobaoSubwayCrowdOfflineLayeredfindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubwayCrowdOfflineLayeredfindRequest()
	},
}

// GetTaobaoSubwayCrowdOfflineLayeredfindRequest 从 sync.Pool 获取 TaobaoSubwayCrowdOfflineLayeredfindAPIRequest
func GetTaobaoSubwayCrowdOfflineLayeredfindAPIRequest() *TaobaoSubwayCrowdOfflineLayeredfindAPIRequest {
	return poolTaobaoSubwayCrowdOfflineLayeredfindAPIRequest.Get().(*TaobaoSubwayCrowdOfflineLayeredfindAPIRequest)
}

// ReleaseTaobaoSubwayCrowdOfflineLayeredfindAPIRequest 将 TaobaoSubwayCrowdOfflineLayeredfindAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubwayCrowdOfflineLayeredfindAPIRequest(v *TaobaoSubwayCrowdOfflineLayeredfindAPIRequest) {
	v.Reset()
	poolTaobaoSubwayCrowdOfflineLayeredfindAPIRequest.Put(v)
}
