package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecentercontractssearchAPIRequest 获取合同类的服务工单信息 API请求
// tmall.servicecenter.contracts.search
//
// 获取合同类的服务工单信息
type TmallservicecentercontractssearchAPIRequest struct {
	model.Params
	// 开始时间:  开始时间和结束时间不能超过15分钟
	_start int64
	// 结束时间:  开始时间和结束时间不能超过15分钟
	_end int64
}

// NewTmallservicecentercontractssearchRequest 初始化TmallservicecentercontractssearchAPIRequest对象
func NewTmallservicecentercontractssearchRequest() *TmallservicecentercontractssearchAPIRequest {
	return &TmallservicecentercontractssearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecentercontractssearchAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.contracts.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecentercontractssearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecentercontractssearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStart is Start Setter
// 开始时间:  开始时间和结束时间不能超过15分钟
func (r *TmallservicecentercontractssearchAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TmallservicecentercontractssearchAPIRequest) GetStart() int64 {
	return r._start
}

// SetEnd is End Setter
// 结束时间:  开始时间和结束时间不能超过15分钟
func (r *TmallservicecentercontractssearchAPIRequest) SetEnd(_end int64) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// GetEnd End Getter
func (r TmallservicecentercontractssearchAPIRequest) GetEnd() int64 {
	return r._end
}
