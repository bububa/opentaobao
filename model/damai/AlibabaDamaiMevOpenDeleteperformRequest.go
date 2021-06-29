package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-场次接口deletePerform APIRequest
alibaba.damai.mev.open.deleteperform

deletePerform
*/
type AlibabaDamaiMevOpenDeleteperformRequest struct {
    model.Params

    // 入参deletePerformParam
    deletePerformParam   *PerformIdOpenParam 

}

func NewAlibabaDamaiMevOpenDeleteperformRequest() *AlibabaDamaiMevOpenDeleteperformRequest{
    return &AlibabaDamaiMevOpenDeleteperformRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenDeleteperformRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.deleteperform"
}

func (r AlibabaDamaiMevOpenDeleteperformRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenDeleteperformRequest) SetDeletePerformParam(deletePerformParam *PerformIdOpenParam) error {
    r.deletePerformParam = deletePerformParam
    r.Set("delete_perform_param", deletePerformParam)
    return nil
}

func (r AlibabaDamaiMevOpenDeleteperformRequest) GetDeletePerformParam() *PerformIdOpenParam {
    return r.deletePerformParam
}

