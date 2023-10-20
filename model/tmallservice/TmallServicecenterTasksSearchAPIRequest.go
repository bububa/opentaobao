package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecentertaskssearchAPIRequest 查询任务类工单信息 API请求
// tmall.servicecenter.tasks.search
//
// 查询任务类工单信息
type TmallservicecentertaskssearchAPIRequest struct {
	model.Params
	// 开始时间:  开始时间和结束时间不能超过15分钟
	_start int64
	// 结束时间:  开始时间和结束时间不能超过15分钟
	_end int64
}

// NewTmallservicecentertaskssearchRequest 初始化TmallservicecentertaskssearchAPIRequest对象
func NewTmallservicecentertaskssearchRequest() *TmallservicecentertaskssearchAPIRequest {
	return &TmallservicecentertaskssearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecentertaskssearchAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.tasks.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecentertaskssearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecentertaskssearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStart is Start Setter
// 开始时间:  开始时间和结束时间不能超过15分钟
func (r *TmallservicecentertaskssearchAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TmallservicecentertaskssearchAPIRequest) GetStart() int64 {
	return r._start
}

// SetEnd is End Setter
// 结束时间:  开始时间和结束时间不能超过15分钟
func (r *TmallservicecentertaskssearchAPIRequest) SetEnd(_end int64) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// GetEnd End Getter
func (r TmallservicecentertaskssearchAPIRequest) GetEnd() int64 {
	return r._end
}
