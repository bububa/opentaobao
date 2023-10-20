package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupidsChangedGetAPIRequest 获取修改的推广组ID API请求
// taobao.simba.adgroupids.changed.get
//
// 获取修改的推广组ID
type TaobaoSimbaAdgroupidsChangedGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 得到此时间点之后的数据，不能大于一个月
	_startTime string
	// 返回的每页数据量大小,默认200最大1000
	_pageSize int64
	// 返回的第几页数据，默认为1
	_pageNo int64
}

// NewTaobaoSimbaAdgroupidsChangedGetRequest 初始化TaobaoSimbaAdgroupidsChangedGetAPIRequest对象
func NewTaobaoSimbaAdgroupidsChangedGetRequest() *TaobaoSimbaAdgroupidsChangedGetAPIRequest {
	return &TaobaoSimbaAdgroupidsChangedGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaAdgroupidsChangedGetAPIRequest) Reset() {
	r._nick = ""
	r._startTime = ""
	r._pageSize = 0
	r._pageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupidsChangedGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroupids.changed.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupidsChangedGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaAdgroupidsChangedGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupidsChangedGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaAdgroupidsChangedGetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 得到此时间点之后的数据，不能大于一个月
func (r *TaobaoSimbaAdgroupidsChangedGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSimbaAdgroupidsChangedGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetPageSize is PageSize Setter
// 返回的每页数据量大小,默认200最大1000
func (r *TaobaoSimbaAdgroupidsChangedGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSimbaAdgroupidsChangedGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaoSimbaAdgroupidsChangedGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoSimbaAdgroupidsChangedGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

var poolTaobaoSimbaAdgroupidsChangedGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaAdgroupidsChangedGetRequest()
	},
}

// GetTaobaoSimbaAdgroupidsChangedGetRequest 从 sync.Pool 获取 TaobaoSimbaAdgroupidsChangedGetAPIRequest
func GetTaobaoSimbaAdgroupidsChangedGetAPIRequest() *TaobaoSimbaAdgroupidsChangedGetAPIRequest {
	return poolTaobaoSimbaAdgroupidsChangedGetAPIRequest.Get().(*TaobaoSimbaAdgroupidsChangedGetAPIRequest)
}

// ReleaseTaobaoSimbaAdgroupidsChangedGetAPIRequest 将 TaobaoSimbaAdgroupidsChangedGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaAdgroupidsChangedGetAPIRequest(v *TaobaoSimbaAdgroupidsChangedGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaAdgroupidsChangedGetAPIRequest.Put(v)
}
