package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayAccountOfflineLayeredfindAPIRequest 获取账户历史报表30天转化周期 API请求
// taobao.subway.account.offline.layeredfind
//
// 获取账户历史报表
type TaobaoSubwayAccountOfflineLayeredfindAPIRequest struct {
	model.Params
	// 开始时间
	_startTime string
	// 结束时间
	_endTime string
	// 页码（0为第一页）
	_offset int64
	// 每页显示的记录条数
	_pageSize int64
	// 转化周期30-30天
	_effect int64
}

// NewTaobaoSubwayAccountOfflineLayeredfindRequest 初始化TaobaoSubwayAccountOfflineLayeredfindAPIRequest对象
func NewTaobaoSubwayAccountOfflineLayeredfindRequest() *TaobaoSubwayAccountOfflineLayeredfindAPIRequest {
	return &TaobaoSubwayAccountOfflineLayeredfindAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubwayAccountOfflineLayeredfindAPIRequest) Reset() {
	r._startTime = ""
	r._endTime = ""
	r._offset = 0
	r._pageSize = 0
	r._effect = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayAccountOfflineLayeredfindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.account.offline.layeredfind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayAccountOfflineLayeredfindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayAccountOfflineLayeredfindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartTime is StartTime Setter
// 开始时间
func (r *TaobaoSubwayAccountOfflineLayeredfindAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSubwayAccountOfflineLayeredfindAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 结束时间
func (r *TaobaoSubwayAccountOfflineLayeredfindAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaoSubwayAccountOfflineLayeredfindAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetOffset is Offset Setter
// 页码（0为第一页）
func (r *TaobaoSubwayAccountOfflineLayeredfindAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// GetOffset Offset Getter
func (r TaobaoSubwayAccountOfflineLayeredfindAPIRequest) GetOffset() int64 {
	return r._offset
}

// SetPageSize is PageSize Setter
// 每页显示的记录条数
func (r *TaobaoSubwayAccountOfflineLayeredfindAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSubwayAccountOfflineLayeredfindAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetEffect is Effect Setter
// 转化周期30-30天
func (r *TaobaoSubwayAccountOfflineLayeredfindAPIRequest) SetEffect(_effect int64) error {
	r._effect = _effect
	r.Set("effect", _effect)
	return nil
}

// GetEffect Effect Getter
func (r TaobaoSubwayAccountOfflineLayeredfindAPIRequest) GetEffect() int64 {
	return r._effect
}

var poolTaobaoSubwayAccountOfflineLayeredfindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubwayAccountOfflineLayeredfindRequest()
	},
}

// GetTaobaoSubwayAccountOfflineLayeredfindRequest 从 sync.Pool 获取 TaobaoSubwayAccountOfflineLayeredfindAPIRequest
func GetTaobaoSubwayAccountOfflineLayeredfindAPIRequest() *TaobaoSubwayAccountOfflineLayeredfindAPIRequest {
	return poolTaobaoSubwayAccountOfflineLayeredfindAPIRequest.Get().(*TaobaoSubwayAccountOfflineLayeredfindAPIRequest)
}

// ReleaseTaobaoSubwayAccountOfflineLayeredfindAPIRequest 将 TaobaoSubwayAccountOfflineLayeredfindAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubwayAccountOfflineLayeredfindAPIRequest(v *TaobaoSubwayAccountOfflineLayeredfindAPIRequest) {
	v.Reset()
	poolTaobaoSubwayAccountOfflineLayeredfindAPIRequest.Put(v)
}
