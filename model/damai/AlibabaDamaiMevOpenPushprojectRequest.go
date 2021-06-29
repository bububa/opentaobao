package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-项目接口pushProject APIRequest
alibaba.damai.mev.open.pushproject

pushProject
*/
type AlibabaDamaiMevOpenPushprojectRequest struct {
    model.Params

    // 入参pushProjectParam
    pushProjectParam   *ThirdProjectPushOpenParam 

}

func NewAlibabaDamaiMevOpenPushprojectRequest() *AlibabaDamaiMevOpenPushprojectRequest{
    return &AlibabaDamaiMevOpenPushprojectRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenPushprojectRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.pushproject"
}

func (r AlibabaDamaiMevOpenPushprojectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenPushprojectRequest) SetPushProjectParam(pushProjectParam *ThirdProjectPushOpenParam) error {
    r.pushProjectParam = pushProjectParam
    r.Set("push_project_param", pushProjectParam)
    return nil
}

func (r AlibabaDamaiMevOpenPushprojectRequest) GetPushProjectParam() *ThirdProjectPushOpenParam {
    return r.pushProjectParam
}

