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
type AlibabaLegalSuitFileUploadAPIRequest struct {
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

// 初始化AlibabaLegalSuitFileUploadAPIRequest对象
func NewAlibabaLegalSuitFileUploadRequest() *AlibabaLegalSuitFileUploadAPIRequest{
    return &AlibabaLegalSuitFileUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitFileUploadAPIRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.file.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitFileUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// File Setter
// 文件
func (r *AlibabaLegalSuitFileUploadAPIRequest) SetFile(_file *model.File) error {
    r._file = _file
    r.Set("file", _file)
    return nil
}

// File Getter
func (r AlibabaLegalSuitFileUploadAPIRequest) GetFile() *model.File {
    return r._file
}
// TimeStamp Setter
// 时间搓
func (r *AlibabaLegalSuitFileUploadAPIRequest) SetTimeStamp(_timeStamp int64) error {
    r._timeStamp = _timeStamp
    r.Set("time_stamp", _timeStamp)
    return nil
}

// TimeStamp Getter
func (r AlibabaLegalSuitFileUploadAPIRequest) GetTimeStamp() int64 {
    return r._timeStamp
}
// FileName Setter
// 文件名称
func (r *AlibabaLegalSuitFileUploadAPIRequest) SetFileName(_fileName string) error {
    r._fileName = _fileName
    r.Set("file_name", _fileName)
    return nil
}

// FileName Getter
func (r AlibabaLegalSuitFileUploadAPIRequest) GetFileName() string {
    return r._fileName
}
// FileSize Setter
// 文件大小
func (r *AlibabaLegalSuitFileUploadAPIRequest) SetFileSize(_fileSize int64) error {
    r._fileSize = _fileSize
    r.Set("file_size", _fileSize)
    return nil
}

// FileSize Getter
func (r AlibabaLegalSuitFileUploadAPIRequest) GetFileSize() int64 {
    return r._fileSize
}
// Signature Setter
// 签名
func (r *AlibabaLegalSuitFileUploadAPIRequest) SetSignature(_signature string) error {
    r._signature = _signature
    r.Set("signature", _signature)
    return nil
}

// Signature Getter
func (r AlibabaLegalSuitFileUploadAPIRequest) GetSignature() string {
    return r._signature
}
