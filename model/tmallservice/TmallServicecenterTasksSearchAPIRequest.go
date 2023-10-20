package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterTasksSearchAPIRequest 查询任务类工单信息 API请求
// tmall.servicecenter.tasks.search
//
// 查询任务类工单信息
type TmallServicecenterTasksSearchAPIRequest struct {
	model.Params
	// 开始时间:  开始时间和结束时间不能超过15分钟
	_start int64
	// 结束时间:  开始时间和结束时间不能超过15分钟
	_end int64
}

// NewTmallServicecenterTasksSearchRequest 初始化TmallServicecenterTasksSearchAPIRequest对象
func NewTmallServicecenterTasksSearchRequest() *TmallServicecenterTasksSearchAPIRequest {
	return &TmallServicecenterTasksSearchAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterTasksSearchAPIRequest) Reset() {
	r._start = 0
	r._end = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterTasksSearchAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.tasks.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterTasksSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterTasksSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStart is Start Setter
// 开始时间:  开始时间和结束时间不能超过15分钟
func (r *TmallServicecenterTasksSearchAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TmallServicecenterTasksSearchAPIRequest) GetStart() int64 {
	return r._start
}

// SetEnd is End Setter
// 结束时间:  开始时间和结束时间不能超过15分钟
func (r *TmallServicecenterTasksSearchAPIRequest) SetEnd(_end int64) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// GetEnd End Getter
func (r TmallServicecenterTasksSearchAPIRequest) GetEnd() int64 {
	return r._end
}

var poolTmallServicecenterTasksSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterTasksSearchRequest()
	},
}

// GetTmallServicecenterTasksSearchRequest 从 sync.Pool 获取 TmallServicecenterTasksSearchAPIRequest
func GetTmallServicecenterTasksSearchAPIRequest() *TmallServicecenterTasksSearchAPIRequest {
	return poolTmallServicecenterTasksSearchAPIRequest.Get().(*TmallServicecenterTasksSearchAPIRequest)
}

// ReleaseTmallServicecenterTasksSearchAPIRequest 将 TmallServicecenterTasksSearchAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterTasksSearchAPIRequest(v *TmallServicecenterTasksSearchAPIRequest) {
	v.Reset()
	poolTmallServicecenterTasksSearchAPIRequest.Put(v)
}
