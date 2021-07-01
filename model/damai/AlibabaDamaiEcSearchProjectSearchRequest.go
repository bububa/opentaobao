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
type AlibabaDamaiEcSearchProjectSearchAPIRequest struct {
    model.Params
    // 入参对象
    _param   *TopSearchProjectParam
}

// 初始化AlibabaDamaiEcSearchProjectSearchAPIRequest对象
func NewAlibabaDamaiEcSearchProjectSearchRequest() *AlibabaDamaiEcSearchProjectSearchAPIRequest{
    return &AlibabaDamaiEcSearchProjectSearchAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiEcSearchProjectSearchAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.ec.search.project.search"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiEcSearchProjectSearchAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参对象
func (r *AlibabaDamaiEcSearchProjectSearchAPIRequest) SetParam(_param *TopSearchProjectParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaDamaiEcSearchProjectSearchAPIRequest) GetParam() *TopSearchProjectParam {
    return r._param
}
