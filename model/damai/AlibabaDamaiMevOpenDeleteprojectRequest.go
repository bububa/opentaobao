package damai

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大麦换验平台-第三方对外开放-项目接口deleteProject API请求
alibaba.damai.mev.open.deleteproject

deleteProject
*/
type AlibabaDamaiMevOpenDeleteprojectRequest struct {
    model.Params
    // 入参deleteProjectParam
    deleteProjectParam   *ProjectIdOpenParam
}

// 初始化AlibabaDamaiMevOpenDeleteprojectRequest对象
func NewAlibabaDamaiMevOpenDeleteprojectRequest() *AlibabaDamaiMevOpenDeleteprojectRequest{
    return &AlibabaDamaiMevOpenDeleteprojectRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeleteprojectRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.deleteproject"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeleteprojectRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeleteProjectParam Setter
// 入参deleteProjectParam
func (r *AlibabaDamaiMevOpenDeleteprojectRequest) SetDeleteProjectParam(deleteProjectParam *ProjectIdOpenParam) error {
    r.deleteProjectParam = deleteProjectParam
    r.Set("delete_project_param", deleteProjectParam)
    return nil
}

// DeleteProjectParam Getter
func (r AlibabaDamaiMevOpenDeleteprojectRequest) GetDeleteProjectParam() *ProjectIdOpenParam {
    return r.deleteProjectParam
}
