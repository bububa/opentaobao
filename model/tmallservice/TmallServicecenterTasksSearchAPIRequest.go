package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallServicecenterTasksSearchAPIRequest
查询任务类工单信息 API请求
tmall.servicecenter.tasks.search

查询任务类工单信息 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterTasksSearchAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.tasks.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterTasksSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Start Setter
// 开始时间:  开始时间和结束时间不能超过15分钟
func (r *TmallServicecenterTasksSearchAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// Get Start Getter
func (r TmallServicecenterTasksSearchAPIRequest) GetStart() int64 {
	return r._start
}

// Set is End Setter
// 结束时间:  开始时间和结束时间不能超过15分钟
func (r *TmallServicecenterTasksSearchAPIRequest) SetEnd(_end int64) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// Get End Getter
func (r TmallServicecenterTasksSearchAPIRequest) GetEnd() int64 {
	return r._end
}
