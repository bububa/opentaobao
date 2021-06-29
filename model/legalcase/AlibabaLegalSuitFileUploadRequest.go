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
    _file   *model.File
    // 时间搓
    _timeStamp   int64
    // 文件名称
    _fileName   string
    // 文件大小
    _fileSize   int64
    // 签名
    _signature   string
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
func (r *AlibabaLegalSuitFileUploadRequest) SetFile(_file *model.File) error {
    r._file = _file
    r.Set("file", _file)
    return nil
}

// File Getter
func (r AlibabaLegalSuitFileUploadRequest) GetFile() *model.File {
    return r._file
}
// TimeStamp Setter
// 时间搓
func (r *AlibabaLegalSuitFileUploadRequest) SetTimeStamp(_timeStamp int64) error {
    r._timeStamp = _timeStamp
    r.Set("time_stamp", _timeStamp)
    return nil
}

// TimeStamp Getter
func (r AlibabaLegalSuitFileUploadRequest) GetTimeStamp() int64 {
    return r._timeStamp
}
// FileName Setter
// 文件名称
func (r *AlibabaLegalSuitFileUploadRequest) SetFileName(_fileName string) error {
    r._fileName = _fileName
    r.Set("file_name", _fileName)
    return nil
}

// FileName Getter
func (r AlibabaLegalSuitFileUploadRequest) GetFileName() string {
    return r._fileName
}
// FileSize Setter
// 文件大小
func (r *AlibabaLegalSuitFileUploadRequest) SetFileSize(_fileSize int64) error {
    r._fileSize = _fileSize
    r.Set("file_size", _fileSize)
    return nil
}

// FileSize Getter
func (r AlibabaLegalSuitFileUploadRequest) GetFileSize() int64 {
    return r._fileSize
}
// Signature Setter
// 签名
func (r *AlibabaLegalSuitFileUploadRequest) SetSignature(_signature string) error {
    r._signature = _signature
    r.Set("signature", _signature)
    return nil
}

// Signature Getter
func (r AlibabaLegalSuitFileUploadRequest) GetSignature() string {
    return r._signature
}
