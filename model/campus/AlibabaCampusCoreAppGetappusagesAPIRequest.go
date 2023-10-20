package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampuscoreappgetappusagesAPIRequest 根据应用ID获得应用所在的园区 API请求
// alibaba.campus.core.app.getappusages
//
// 传入应用的id,  获得用户授权的园区
type AlibabacampuscoreappgetappusagesAPIRequest struct {
	model.Params
	// WorkBenchContext
	_workBenchContext *WorkBenchContext
	// 应用id
	_appid int64
}

// NewAlibabacampuscoreappgetappusagesRequest 初始化AlibabacampuscoreappgetappusagesAPIRequest对象
func NewAlibabacampuscoreappgetappusagesRequest() *AlibabacampuscoreappgetappusagesAPIRequest {
	return &AlibabacampuscoreappgetappusagesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacampuscoreappgetappusagesAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.core.app.getappusages"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacampuscoreappgetappusagesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacampuscoreappgetappusagesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkBenchContext is WorkBenchContext Setter
// WorkBenchContext
func (r *AlibabacampuscoreappgetappusagesAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabacampuscoreappgetappusagesAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// SetAppid is Appid Setter
// 应用id
func (r *AlibabacampuscoreappgetappusagesAPIRequest) SetAppid(_appid int64) error {
	r._appid = _appid
	r.Set("appid", _appid)
	return nil
}

// GetAppid Appid Getter
func (r AlibabacampuscoreappgetappusagesAPIRequest) GetAppid() int64 {
	return r._appid
}
