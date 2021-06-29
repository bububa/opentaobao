package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
文件上传API APIRequest
yunos.tvpubadmin.common.file.upload

文件上传服务
*/
type YunosTvpubadminCommonFileUploadRequest struct {
    model.Params

    // 文件字节流
    bytes   []*model.File 

    // 原文件名
    originalFilename   string 

    // 文件大小
    size   string 

    // 文件类型
    contentType   string 

    // 上传地址
    uploadPath   string 

}

func NewYunosTvpubadminCommonFileUploadRequest() *YunosTvpubadminCommonFileUploadRequest{
    return &YunosTvpubadminCommonFileUploadRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminCommonFileUploadRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.common.file.upload"
}

func (r YunosTvpubadminCommonFileUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminCommonFileUploadRequest) SetBytes(bytes []*model.File) error {
    r.bytes = bytes
    r.Set("bytes", bytes)
    return nil
}

func (r YunosTvpubadminCommonFileUploadRequest) GetBytes() []*model.File {
    return r.bytes
}

func (r *YunosTvpubadminCommonFileUploadRequest) SetOriginalFilename(originalFilename string) error {
    r.originalFilename = originalFilename
    r.Set("original_filename", originalFilename)
    return nil
}

func (r YunosTvpubadminCommonFileUploadRequest) GetOriginalFilename() string {
    return r.originalFilename
}

func (r *YunosTvpubadminCommonFileUploadRequest) SetSize(size string) error {
    r.size = size
    r.Set("size", size)
    return nil
}

func (r YunosTvpubadminCommonFileUploadRequest) GetSize() string {
    return r.size
}

func (r *YunosTvpubadminCommonFileUploadRequest) SetContentType(contentType string) error {
    r.contentType = contentType
    r.Set("content_type", contentType)
    return nil
}

func (r YunosTvpubadminCommonFileUploadRequest) GetContentType() string {
    return r.contentType
}

func (r *YunosTvpubadminCommonFileUploadRequest) SetUploadPath(uploadPath string) error {
    r.uploadPath = uploadPath
    r.Set("upload_path", uploadPath)
    return nil
}

func (r YunosTvpubadminCommonFileUploadRequest) GetUploadPath() string {
    return r.uploadPath
}

