package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupidsDeletedGetAPIRequest 获取删除的推广组ID API请求
// taobao.simba.adgroupids.deleted.get
//
// 获取删除的推广组ID
type TaobaoSimbaAdgroupidsDeletedGetAPIRequest struct {
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

// NewTaobaoSimbaAdgroupidsDeletedGetRequest 初始化TaobaoSimbaAdgroupidsDeletedGetAPIRequest对象
func NewTaobaoSimbaAdgroupidsDeletedGetRequest() *TaobaoSimbaAdgroupidsDeletedGetAPIRequest {
	return &TaobaoSimbaAdgroupidsDeletedGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaAdgroupidsDeletedGetAPIRequest) Reset() {
	r._nick = ""
	r._startTime = ""
	r._pageSize = 0
	r._pageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupidsDeletedGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroupids.deleted.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupidsDeletedGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaAdgroupidsDeletedGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupidsDeletedGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaAdgroupidsDeletedGetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 得到此时间点之后的数据，不能大于一个月
func (r *TaobaoSimbaAdgroupidsDeletedGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSimbaAdgroupidsDeletedGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetPageSize is PageSize Setter
// 返回的每页数据量大小,默认200最大1000
func (r *TaobaoSimbaAdgroupidsDeletedGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSimbaAdgroupidsDeletedGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaoSimbaAdgroupidsDeletedGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoSimbaAdgroupidsDeletedGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

var poolTaobaoSimbaAdgroupidsDeletedGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaAdgroupidsDeletedGetRequest()
	},
}

// GetTaobaoSimbaAdgroupidsDeletedGetRequest 从 sync.Pool 获取 TaobaoSimbaAdgroupidsDeletedGetAPIRequest
func GetTaobaoSimbaAdgroupidsDeletedGetAPIRequest() *TaobaoSimbaAdgroupidsDeletedGetAPIRequest {
	return poolTaobaoSimbaAdgroupidsDeletedGetAPIRequest.Get().(*TaobaoSimbaAdgroupidsDeletedGetAPIRequest)
}

// ReleaseTaobaoSimbaAdgroupidsDeletedGetAPIRequest 将 TaobaoSimbaAdgroupidsDeletedGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaAdgroupidsDeletedGetAPIRequest(v *TaobaoSimbaAdgroupidsDeletedGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaAdgroupidsDeletedGetAPIRequest.Put(v)
}
