package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交KA合作楼盘 API请求
alibaba.alihouse.newhome.project.cooperation.submit

提交KA合作楼盘
*/
type AlibabaAlihouseNewhomeProjectCooperationSubmitRequest struct {
    model.Params
    // ka合作对象
    _projectCooperationDto   *ProjectCooperationDTO
}

// 初始化AlibabaAlihouseNewhomeProjectCooperationSubmitRequest对象
func NewAlibabaAlihouseNewhomeProjectCooperationSubmitRequest() *AlibabaAlihouseNewhomeProjectCooperationSubmitRequest{
    return &AlibabaAlihouseNewhomeProjectCooperationSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectCooperationSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.cooperation.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectCooperationSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProjectCooperationDto Setter
// ka合作对象
func (r *AlibabaAlihouseNewhomeProjectCooperationSubmitRequest) SetProjectCooperationDto(_projectCooperationDto *ProjectCooperationDTO) error {
    r._projectCooperationDto = _projectCooperationDto
    r.Set("project_cooperation_dto", _projectCooperationDto)
    return nil
}

// ProjectCooperationDto Getter
func (r AlibabaAlihouseNewhomeProjectCooperationSubmitRequest) GetProjectCooperationDto() *ProjectCooperationDTO {
    return r._projectCooperationDto
}
