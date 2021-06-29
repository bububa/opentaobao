package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-项目接口deleteProject APIRequest
alibaba.damai.mev.open.deleteproject

deleteProject
*/
type AlibabaDamaiMevOpenDeleteprojectRequest struct {
    model.Params

    // 入参deleteProjectParam
    deleteProjectParam   *ProjectIdOpenParam 

}

func NewAlibabaDamaiMevOpenDeleteprojectRequest() *AlibabaDamaiMevOpenDeleteprojectRequest{
    return &AlibabaDamaiMevOpenDeleteprojectRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDamaiMevOpenDeleteprojectRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.deleteproject"
}

func (r AlibabaDamaiMevOpenDeleteprojectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDamaiMevOpenDeleteprojectRequest) SetDeleteProjectParam(deleteProjectParam *ProjectIdOpenParam) error {
    r.deleteProjectParam = deleteProjectParam
    r.Set("delete_project_param", deleteProjectParam)
    return nil
}

func (r AlibabaDamaiMevOpenDeleteprojectRequest) GetDeleteProjectParam() *ProjectIdOpenParam {
    return r.deleteProjectParam
}

