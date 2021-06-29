package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
诉讼文件上传接口 API请求
alibaba.legal.suit.file.upload

上传文件接口
*/
type AlibabaLegalSuitFileUploadRequest struct {
    model.Params
    // 文件
    file   []*model.File
    // 时间搓
    timeStamp   int64
    // 文件名称
    fileName   string
    // 文件大小
    fileSize   int64
    // 签名
    signature   string
}

// 初始化AlibabaLegalSuitFileUploadRequest对象
func NewAlibabaLegalSuitFileUploadRequest() *AlibabaLegalSuitFileUploadRequest{
    return &AlibabaLegalSuitFileUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitFileUploadRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.file.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitFileUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// File Setter
// 文件
func (r *AlibabaLegalSuitFileUploadRequest) SetFile(file []*model.File) error {
    r.file = file
    r.Set("file", file)
    return nil
}

// File Getter
func (r AlibabaLegalSuitFileUploadRequest) GetFile() []*model.File {
    return r.file
}
// TimeStamp Setter
// 时间搓
func (r *AlibabaLegalSuitFileUploadRequest) SetTimeStamp(timeStamp int64) error {
    r.timeStamp = timeStamp
    r.Set("time_stamp", timeStamp)
    return nil
}

// TimeStamp Getter
func (r AlibabaLegalSuitFileUploadRequest) GetTimeStamp() int64 {
    return r.timeStamp
}
// FileName Setter
// 文件名称
func (r *AlibabaLegalSuitFileUploadRequest) SetFileName(fileName string) error {
    r.fileName = fileName
    r.Set("file_name", fileName)
    return nil
}

// FileName Getter
func (r AlibabaLegalSuitFileUploadRequest) GetFileName() string {
    return r.fileName
}
// FileSize Setter
// 文件大小
func (r *AlibabaLegalSuitFileUploadRequest) SetFileSize(fileSize int64) error {
    r.fileSize = fileSize
    r.Set("file_size", fileSize)
    return nil
}

// FileSize Getter
func (r AlibabaLegalSuitFileUploadRequest) GetFileSize() int64 {
    return r.fileSize
}
// Signature Setter
// 签名
func (r *AlibabaLegalSuitFileUploadRequest) SetSignature(signature string) error {
    r.signature = signature
    r.Set("signature", signature)
    return nil
}

// Signature Getter
func (r AlibabaLegalSuitFileUploadRequest) GetSignature() string {
    return r.signature
}
