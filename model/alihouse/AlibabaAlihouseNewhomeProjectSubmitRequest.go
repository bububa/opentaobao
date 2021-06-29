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
type AlibabaAlihouseNewhomeProjectSubmitRequest struct {
    model.Params
    // 楼盘对象
    ebbasProjectDto   *EbbasProjectDto
}

// 初始化AlibabaAlihouseNewhomeProjectSubmitRequest对象
func NewAlibabaAlihouseNewhomeProjectSubmitRequest() *AlibabaAlihouseNewhomeProjectSubmitRequest{
    return &AlibabaAlihouseNewhomeProjectSubmitRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectSubmitRequest) GetApiMethodName() string {
    return "alibaba.alihouse.newhome.project.submit"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectSubmitRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EbbasProjectDto Setter
// 楼盘对象
func (r *AlibabaAlihouseNewhomeProjectSubmitRequest) SetEbbasProjectDto(ebbasProjectDto *EbbasProjectDto) error {
    r.ebbasProjectDto = ebbasProjectDto
    r.Set("ebbas_project_dto", ebbasProjectDto)
    return nil
}

// EbbasProjectDto Getter
func (r AlibabaAlihouseNewhomeProjectSubmitRequest) GetEbbasProjectDto() *EbbasProjectDto {
    return r.ebbasProjectDto
}
