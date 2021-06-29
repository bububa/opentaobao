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
type AlibabaCampusCoreAppGetappusagesRequest struct {
    model.Params
    // 应用id
    appid   int64
    // WorkBenchContext
    workBenchContext   *WorkBenchContext
}

// 初始化AlibabaCampusCoreAppGetappusagesRequest对象
func NewAlibabaCampusCoreAppGetappusagesRequest() *AlibabaCampusCoreAppGetappusagesRequest{
    return &AlibabaCampusCoreAppGetappusagesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCampusCoreAppGetappusagesRequest) GetApiMethodName() string {
    return "alibaba.campus.core.app.getappusages"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCampusCoreAppGetappusagesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Appid Setter
// 应用id
func (r *AlibabaCampusCoreAppGetappusagesRequest) SetAppid(appid int64) error {
    r.appid = appid
    r.Set("appid", appid)
    return nil
}

// Appid Getter
func (r AlibabaCampusCoreAppGetappusagesRequest) GetAppid() int64 {
    return r.appid
}
// WorkBenchContext Setter
// WorkBenchContext
func (r *AlibabaCampusCoreAppGetappusagesRequest) SetWorkBenchContext(workBenchContext *WorkBenchContext) error {
    r.workBenchContext = workBenchContext
    r.Set("work_bench_context", workBenchContext)
    return nil
}

// WorkBenchContext Getter
func (r AlibabaCampusCoreAppGetappusagesRequest) GetWorkBenchContext() *WorkBenchContext {
    return r.workBenchContext
}
