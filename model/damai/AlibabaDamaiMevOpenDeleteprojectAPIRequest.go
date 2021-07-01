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
type AlibabaDamaiMevOpenDeleteprojectAPIRequest struct {
    model.Params
    // 入参deleteProjectParam
    _deleteProjectParam   *ProjectIdOpenParam
}

// 初始化AlibabaDamaiMevOpenDeleteprojectAPIRequest对象
func NewAlibabaDamaiMevOpenDeleteprojectRequest() *AlibabaDamaiMevOpenDeleteprojectAPIRequest{
    return &AlibabaDamaiMevOpenDeleteprojectAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMevOpenDeleteprojectAPIRequest) GetApiMethodName() string {
    return "alibaba.damai.mev.open.deleteproject"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMevOpenDeleteprojectAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DeleteProjectParam Setter
// 入参deleteProjectParam
func (r *AlibabaDamaiMevOpenDeleteprojectAPIRequest) SetDeleteProjectParam(_deleteProjectParam *ProjectIdOpenParam) error {
    r._deleteProjectParam = _deleteProjectParam
    r.Set("delete_project_param", _deleteProjectParam)
    return nil
}

// DeleteProjectParam Getter
func (r AlibabaDamaiMevOpenDeleteprojectAPIRequest) GetDeleteProjectParam() *ProjectIdOpenParam {
    return r._deleteProjectParam
}
