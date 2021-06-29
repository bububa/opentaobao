package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交楼盘电话 API请求
alibaba.alihouse.newhome.project.phone.submit

提交楼盘电话
*/
type AlibabaAlihouseNewhomeProjectPhoneSubmitRequest struct {
    model.Params
    // 楼盘电话
    _projectPhoneDto   *ProjectPhoneDto
}

// 初始化AlibabaAlihouseNewhomeProjectPhoneSubmitRequest对象
func NewAlibabaAlihouseNewhomeProjectPhoneSubmitRequest() *AlibabaAlihouseNewhomeProjectPhoneSubmitRequest{
    return &AlibabaAlihouseNewhomeProjectPhoneSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectPhoneSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.phone.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectPhoneSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ProjectPhoneDto Setter
// 楼盘电话
func (r *AlibabaAlihouseNewhomeProjectPhoneSubmitRequest) SetProjectPhoneDto(_projectPhoneDto *ProjectPhoneDto) error {
    r._projectPhoneDto = _projectPhoneDto
    r.Set("project_phone_dto", _projectPhoneDto)
    return nil
}

// ProjectPhoneDto Getter
func (r AlibabaAlihouseNewhomeProjectPhoneSubmitRequest) GetProjectPhoneDto() *ProjectPhoneDto {
    return r._projectPhoneDto
}
