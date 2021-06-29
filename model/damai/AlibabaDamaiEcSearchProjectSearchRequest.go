package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦电商对外搜索服务 APIRequest
alibaba.damai.ec.search.project.search

大麦电商对外搜索服务
*/
type AlibabaDamaiEcSearchProjectSearchRequest struct {
    model.Params

    // 入参对象
    param   *TopSearchProjectParam 

}

func NewAlibabaDamaiEcSearchProjectSearchRequest() *AlibabaDamaiEcSearchProjectSearchRequest{
    return &AlibabaDamaiEcSearchProjectSearchRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiEcSearchProjectSearchRequest) GetApiMethodName() string {
    return "alibaba.damai.ec.search.project.search"
}

func (r AlibabaDamaiEcSearchProjectSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiEcSearchProjectSearchRequest) SetParam(param *TopSearchProjectParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaDamaiEcSearchProjectSearchRequest) GetParam() *TopSearchProjectParam {
    return r.param
}

