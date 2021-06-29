package fpm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
结算单文件上传 APIRequest
alibaba.fpm.file.upload

结算单文件上传
*/
type AlibabaFpmFileUploadRequest struct {
    model.Params

    // 实体
    bizDto   *FileUploadRequestDto 

}

func NewAlibabaFpmFileUploadRequest() *AlibabaFpmFileUploadRequest{
    return &AlibabaFpmFileUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFpmFileUploadRequest) GetApiMethodName() string {
    return "alibaba.fpm.file.upload"
}

func (r AlibabaFpmFileUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFpmFileUploadRequest) SetBizDto(bizDto *FileUploadRequestDto) error {
    r.bizDto = bizDto
    r.Set("biz_dto", bizDto)
    return nil
}

func (r AlibabaFpmFileUploadRequest) GetBizDto() *FileUploadRequestDto {
    return r.bizDto
}

