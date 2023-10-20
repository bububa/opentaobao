package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayKeywordOfflineLayeredfindAPIRequest 查询关键词离线报表30天转化周期 API请求
// taobao.subway.keyword.offline.layeredfind
//
// 获取关键词离线报表
type TaobaoSubwayKeywordOfflineLayeredfindAPIRequest struct {
	model.Params
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
	// 数据来源（pc站内：1；pc站外：2；无限站内：4；无限站内：5；销量明星：6）
	_pvTypeIn int64
	// 需要查询的关键词id，不传表示不限
	_bidwordIdEqual int64
	// 页码（0为第一页）
	_offset int64
	// 每页显示的记录条数
	_pageSize int64
	// 转化周期30-30天
	_effect int64
}

// NewTaobaoSubwayKeywordOfflineLayeredfindRequest 初始化TaobaoSubwayKeywordOfflineLayeredfindAPIRequest对象
func NewTaobaoSubwayKeywordOfflineLayeredfindRequest() *TaobaoSubwayKeywordOfflineLayeredfindAPIRequest {
	return &TaobaoSubwayKeywordOfflineLayeredfindAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) Reset() {
	r._startTime = ""
	r._endTime = ""
	r._pvTypeIn = 0
	r._bidwordIdEqual = 0
	r._offset = 0
	r._pageSize = 0
	r._effect = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.keyword.offline.layeredfind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetPvTypeIn is PvTypeIn Setter
// 数据来源（pc站内：1；pc站外：2；无限站内：4；无限站内：5；销量明星：6）
func (r *TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) SetPvTypeIn(_pvTypeIn int64) error {
	r._pvTypeIn = _pvTypeIn
	r.Set("pv_type_in", _pvTypeIn)
	return nil
}

// GetPvTypeIn PvTypeIn Getter
func (r TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) GetPvTypeIn() int64 {
	return r._pvTypeIn
}

// SetBidwordIdEqual is BidwordIdEqual Setter
// 需要查询的关键词id，不传表示不限
func (r *TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) SetBidwordIdEqual(_bidwordIdEqual int64) error {
	r._bidwordIdEqual = _bidwordIdEqual
	r.Set("bidword_id_equal", _bidwordIdEqual)
	return nil
}

// GetBidwordIdEqual BidwordIdEqual Getter
func (r TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) GetBidwordIdEqual() int64 {
	return r._bidwordIdEqual
}

// SetOffset is Offset Setter
// 页码（0为第一页）
func (r *TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 每页显示的记录条数
func (r *TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetEffect is Effect Setter
// 转化周期30-30天
func (r *TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) GetEffect() int64 {
	return r._effect
}

var poolTaobaoSubwayKeywordOfflineLayeredfindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubwayKeywordOfflineLayeredfindRequest()
	},
}

// GetTaobaoSubwayKeywordOfflineLayeredfindRequest 从 sync.Pool 获取 TaobaoSubwayKeywordOfflineLayeredfindAPIRequest
func GetTaobaoSubwayKeywordOfflineLayeredfindAPIRequest() *TaobaoSubwayKeywordOfflineLayeredfindAPIRequest {
	return poolTaobaoSubwayKeywordOfflineLayeredfindAPIRequest.Get().(*TaobaoSubwayKeywordOfflineLayeredfindAPIRequest)
}

// ReleaseTaobaoSubwayKeywordOfflineLayeredfindAPIRequest 将 TaobaoSubwayKeywordOfflineLayeredfindAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubwayKeywordOfflineLayeredfindAPIRequest(v *TaobaoSubwayKeywordOfflineLayeredfindAPIRequest) {
	v.Reset()
	poolTaobaoSubwayKeywordOfflineLayeredfindAPIRequest.Put(v)
}
