package campus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据应用ID获得应用所在的园区 API请求
alibaba.campus.core.app.getappusages

传入应用的id,  获得用户授权的园区
*/
type AlibabaCampusCoreAppGetappusagesAPIRequest struct {
    model.Params
    // 应用id
    _appid   int64
    // WorkBenchContext
    _workBenchContext   *WorkBenchContext
}

// 初始化AlibabaCampusCoreAppGetappusagesAPIRequest对象
func NewAlibabaCampusCoreAppGetappusagesRequest() *AlibabaCampusCoreAppGetappusagesAPIRequest{
    return &AlibabaCampusCoreAppGetappusagesAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusCoreAppGetappusagesAPIRequest) GetApiMethodName() string {
    return "alibaba.campus.core.app.getappusages"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusCoreAppGetappusagesAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Appid Setter
// 应用id
func (r *AlibabaCampusCoreAppGetappusagesAPIRequest) SetAppid(_appid int64) error {
    r._appid = _appid
    r.Set("appid", _appid)
    return nil
}

// Appid Getter
func (r AlibabaCampusCoreAppGetappusagesAPIRequest) GetAppid() int64 {
    return r._appid
}
// WorkBenchContext Setter
// WorkBenchContext
func (r *AlibabaCampusCoreAppGetappusagesAPIRequest) SetWorkBenchContext(_workBenchContext *WorkBenchContext) error {
    r._workBenchContext = _workBenchContext
    r.Set("work_bench_context", _workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaCampusCoreAppGetappusagesAPIRequest) GetWorkBenchContext() *WorkBenchContext {
    return r._workBenchContext
}
