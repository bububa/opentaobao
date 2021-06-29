package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
文件上传API API请求
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

// 初始化YunosTvpubadminCommonFileUploadRequest对象
func NewYunosTvpubadminCommonFileUploadRequest() *YunosTvpubadminCommonFileUploadRequest{
    return &YunosTvpubadminCommonFileUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminCommonFileUploadRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.common.file.upload"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminCommonFileUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Bytes Setter
// 文件字节流
func (r *YunosTvpubadminCommonFileUploadRequest) SetBytes(bytes []*model.File) error {
    r.bytes = bytes
    r.Set("bytes", bytes)
    return nil
}

// Bytes Getter
func (r YunosTvpubadminCommonFileUploadRequest) GetBytes() []*model.File {
    return r.bytes
}
// OriginalFilename Setter
// 原文件名
func (r *YunosTvpubadminCommonFileUploadRequest) SetOriginalFilename(originalFilename string) error {
    r.originalFilename = originalFilename
    r.Set("original_filename", originalFilename)
    return nil
}

// OriginalFilename Getter
func (r YunosTvpubadminCommonFileUploadRequest) GetOriginalFilename() string {
    return r.originalFilename
}
// Size Setter
// 文件大小
func (r *YunosTvpubadminCommonFileUploadRequest) SetSize(size string) error {
    r.size = size
    r.Set("size", size)
    return nil
}

// Size Getter
func (r YunosTvpubadminCommonFileUploadRequest) GetSize() string {
    return r.size
}
// ContentType Setter
// 文件类型
func (r *YunosTvpubadminCommonFileUploadRequest) SetContentType(contentType string) error {
    r.contentType = contentType
    r.Set("content_type", contentType)
    return nil
}

// ContentType Getter
func (r YunosTvpubadminCommonFileUploadRequest) GetContentType() string {
    return r.contentType
}
// UploadPath Setter
// 上传地址
func (r *YunosTvpubadminCommonFileUploadRequest) SetUploadPath(uploadPath string) error {
    r.uploadPath = uploadPath
    r.Set("upload_path", uploadPath)
    return nil
}

// UploadPath Getter
func (r YunosTvpubadminCommonFileUploadRequest) GetUploadPath() string {
    return r.uploadPath
}
