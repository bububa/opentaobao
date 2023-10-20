package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupsChangedGetAPIRequest 分页获取修改的推广组ID和修改时间 API请求
// taobao.simba.adgroups.changed.get
//
// 分页获取修改的推广组ID和修改时间
type TaobaoSimbaAdgroupsChangedGetAPIRequest struct {
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

// NewTaobaoSimbaAdgroupsChangedGetRequest 初始化TaobaoSimbaAdgroupsChangedGetAPIRequest对象
func NewTaobaoSimbaAdgroupsChangedGetRequest() *TaobaoSimbaAdgroupsChangedGetAPIRequest {
	return &TaobaoSimbaAdgroupsChangedGetAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaAdgroupsChangedGetAPIRequest) Reset() {
	r._nick = ""
	r._startTime = ""
	r._pageSize = 0
	r._pageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupsChangedGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroups.changed.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupsChangedGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaAdgroupsChangedGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupsChangedGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaAdgroupsChangedGetAPIRequest) GetNick() string {
	return r._nick
}

// SetStartTime is StartTime Setter
// 得到此时间点之后的数据，不能大于一个月
func (r *TaobaoSimbaAdgroupsChangedGetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaoSimbaAdgroupsChangedGetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetPageSize is PageSize Setter
// 返回的每页数据量大小,默认200最大1000
func (r *TaobaoSimbaAdgroupsChangedGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoSimbaAdgroupsChangedGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 返回的第几页数据，默认为1
func (r *TaobaoSimbaAdgroupsChangedGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoSimbaAdgroupsChangedGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

var poolTaobaoSimbaAdgroupsChangedGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaAdgroupsChangedGetRequest()
	},
}

// GetTaobaoSimbaAdgroupsChangedGetRequest 从 sync.Pool 获取 TaobaoSimbaAdgroupsChangedGetAPIRequest
func GetTaobaoSimbaAdgroupsChangedGetAPIRequest() *TaobaoSimbaAdgroupsChangedGetAPIRequest {
	return poolTaobaoSimbaAdgroupsChangedGetAPIRequest.Get().(*TaobaoSimbaAdgroupsChangedGetAPIRequest)
}

// ReleaseTaobaoSimbaAdgroupsChangedGetAPIRequest 将 TaobaoSimbaAdgroupsChangedGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaAdgroupsChangedGetAPIRequest(v *TaobaoSimbaAdgroupsChangedGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaAdgroupsChangedGetAPIRequest.Put(v)
}
