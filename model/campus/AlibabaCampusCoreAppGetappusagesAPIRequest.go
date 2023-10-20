package campus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusCoreAppGetappusagesAPIRequest 根据应用ID获得应用所在的园区 API请求
// alibaba.campus.core.app.getappusages
//
// 传入应用的id,  获得用户授权的园区
type AlibabaCampusCoreAppGetappusagesAPIRequest struct {
	model.Params
	// WorkBenchContext
	_workBenchContext *WorkBenchContext
	// 应用id
	_appid int64
}

// NewAlibabaCampusCoreAppGetappusagesRequest 初始化AlibabaCampusCoreAppGetappusagesAPIRequest对象
func NewAlibabaCampusCoreAppGetappusagesRequest() *AlibabaCampusCoreAppGetappusagesAPIRequest {
	return &AlibabaCampusCoreAppGetappusagesAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCampusCoreAppGetappusagesAPIRequest) Reset() {
	r._workBenchContext = nil
	r._appid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCampusCoreAppGetappusagesAPIRequest) GetApiMethodName() string {
	return "alibaba.campus.core.app.getappusages"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCampusCoreAppGetappusagesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCampusCoreAppGetappusagesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkBenchContext is WorkBenchContext Setter
// WorkBenchContext
func (r *AlibabaCampusCoreAppGetappusagesAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
	r._workBenchContext = _workBenchContext
	r.Set("work_bench_context", _workBenchContext)
	return nil
}

// GetWorkBenchContext WorkBenchContext Getter
func (r AlibabaCampusCoreAppGetappusagesAPIRequest) GetWorkBenchContext() *WorkBenchContext {
	return r._workBenchContext
}

// SetAppid is Appid Setter
// 应用id
func (r *AlibabaCampusCoreAppGetappusagesAPIRequest) SetAppid(_appid int64) error {
	r._appid = _appid
	r.Set("appid", _appid)
	return nil
}

// GetAppid Appid Getter
func (r AlibabaCampusCoreAppGetappusagesAPIRequest) GetAppid() int64 {
	return r._appid
}

var poolAlibabaCampusCoreAppGetappusagesAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCampusCoreAppGetappusagesRequest()
	},
}

// GetAlibabaCampusCoreAppGetappusagesRequest 从 sync.Pool 获取 AlibabaCampusCoreAppGetappusagesAPIRequest
func GetAlibabaCampusCoreAppGetappusagesAPIRequest() *AlibabaCampusCoreAppGetappusagesAPIRequest {
	return poolAlibabaCampusCoreAppGetappusagesAPIRequest.Get().(*AlibabaCampusCoreAppGetappusagesAPIRequest)
}

// ReleaseAlibabaCampusCoreAppGetappusagesAPIRequest 将 AlibabaCampusCoreAppGetappusagesAPIRequest 放入 sync.Pool
func ReleaseAlibabaCampusCoreAppGetappusagesAPIRequest(v *AlibabaCampusCoreAppGetappusagesAPIRequest) {
	v.Reset()
	poolAlibabaCampusCoreAppGetappusagesAPIRequest.Put(v)
}
