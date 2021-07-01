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
type AlibabaFpmFileUploadAPIRequest struct {
    model.Params
    // 实体
    _bizDto   *FileUploadRequestDTO
}

// 初始化AlibabaFpmFileUploadAPIRequest对象
func NewAlibabaFpmFileUploadRequest() *AlibabaFpmFileUploadAPIRequest{
    return &AlibabaFpmFileUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFpmFileUploadAPIRequest) GetApiMethodName() string {
    return "alibaba.fpm.file.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFpmFileUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BizDto Setter
// 实体
func (r *AlibabaFpmFileUploadAPIRequest) SetBizDto(_bizDto *FileUploadRequestDTO) error {
    r._bizDto = _bizDto
    r.Set("biz_dto", _bizDto)
    return nil
}

// BizDto Getter
func (r AlibabaFpmFileUploadAPIRequest) GetBizDto() *FileUploadRequestDTO {
    return r._bizDto
}
