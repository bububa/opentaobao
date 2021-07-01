package alihouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
提交楼盘信息 API请求
alibaba.alihouse.newhome.project.submit

提交楼盘信息
*/
type AlibabaAlihouseNewhomeProjectSubmitAPIRequest struct {
    model.Params
    // 楼盘对象
    _ebbasProjectDto   *EbbasProjectDTO
}

// 初始化AlibabaAlihouseNewhomeProjectSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectSubmitRequest() *AlibabaAlihouseNewhomeProjectSubmitAPIRequest{
    return &AlibabaAlihouseNewhomeProjectSubmitAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectSubmitAPIRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectSubmitAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EbbasProjectDto Setter
// 楼盘对象
func (r *AlibabaAlihouseNewhomeProjectSubmitAPIRequest) SetEbbasProjectDto(_ebbasProjectDto *EbbasProjectDTO) error {
    r._ebbasProjectDto = _ebbasProjectDto
    r.Set("ebbas_project_dto", _ebbasProjectDto)
    return nil
}

// EbbasProjectDto Getter
func (r AlibabaAlihouseNewhomeProjectSubmitAPIRequest) GetEbbasProjectDto() *EbbasProjectDTO {
    return r._ebbasProjectDto
}
