package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteranomalyrecoursesearchAPIRequest 天猫服务平台服务商一键求助单查询 API请求
// tmall.servicecenter.anomalyrecourse.search
//
// 天猫服务平台服务商一键求助单查询
type TmallservicecenteranomalyrecoursesearchAPIRequest struct {
	model.Params
	// 开始时间
	_start int64
	// 结束时间
	_end int64
}

// NewTmallservicecenteranomalyrecoursesearchRequest 初始化TmallservicecenteranomalyrecoursesearchAPIRequest对象
func NewTmallservicecenteranomalyrecoursesearchRequest() *TmallservicecenteranomalyrecoursesearchAPIRequest {
	return &TmallservicecenteranomalyrecoursesearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenteranomalyrecoursesearchAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenteranomalyrecoursesearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenteranomalyrecoursesearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStart is Start Setter
// 开始时间
func (r *TmallservicecenteranomalyrecoursesearchAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TmallservicecenteranomalyrecoursesearchAPIRequest) GetStart() int64 {
	return r._start
}

// SetEnd is End Setter
// 结束时间
func (r *TmallservicecenteranomalyrecoursesearchAPIRequest) SetEnd(_end int64) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// GetEnd End Getter
func (r TmallservicecenteranomalyrecoursesearchAPIRequest) GetEnd() int64 {
	return r._end
}
