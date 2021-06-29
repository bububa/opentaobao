package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦电商对外搜索服务 API请求
alibaba.damai.ec.search.project.search

大麦电商对外搜索服务
*/
type AlibabaDamaiEcSearchProjectSearchRequest struct {
    model.Params
    // 入参对象
    _param   *TopSearchProjectParam
}

// 初始化AlibabaDamaiEcSearchProjectSearchRequest对象
func NewAlibabaDamaiEcSearchProjectSearchRequest() *AlibabaDamaiEcSearchProjectSearchRequest{
    return &AlibabaDamaiEcSearchProjectSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiEcSearchProjectSearchRequest) GetApiMethodName() string {
    return "alibaba.damai.ec.search.project.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiEcSearchProjectSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参对象
func (r *AlibabaDamaiEcSearchProjectSearchRequest) SetParam(_param *TopSearchProjectParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaDamaiEcSearchProjectSearchRequest) GetParam() *TopSearchProjectParam {
    return r._param
}
