package fpm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
结算单文件上传 API请求
alibaba.fpm.file.upload

结算单文件上传
*/
type AlibabaFpmFileUploadRequest struct {
    model.Params
    // 实体
    bizDto   *FileUploadRequestDto
}

// 初始化AlibabaFpmFileUploadRequest对象
func NewAlibabaFpmFileUploadRequest() *AlibabaFpmFileUploadRequest{
    return &AlibabaFpmFileUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFpmFileUploadRequest) GetApiMethodName() string {
    return "alibaba.fpm.file.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFpmFileUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizDto Setter
// 实体
func (r *AlibabaFpmFileUploadRequest) SetBizDto(bizDto *FileUploadRequestDto) error {
    r.bizDto = bizDto
    r.Set("biz_dto", bizDto)
    return nil
}

// BizDto Getter
func (r AlibabaFpmFileUploadRequest) GetBizDto() *FileUploadRequestDto {
    return r.bizDto
}
