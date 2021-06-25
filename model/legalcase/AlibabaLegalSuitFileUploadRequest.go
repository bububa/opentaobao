package legalcase

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
诉讼文件上传接口 APIRequest
alibaba.legal.suit.file.upload

上传文件接口
*/
type AlibabaLegalSuitFileUploadRequest struct {
    model.Params

    // 文件
    file   []byte 

    // 时间搓
    timeStamp   int64 

    // 文件名称
    fileName   string 

    // 文件大小
    fileSize   int64 

    // 签名
    signature   string 

}

func NewAlibabaLegalSuitFileUploadRequest() *AlibabaLegalSuitFileUploadRequest{
    return &AlibabaLegalSuitFileUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLegalSuitFileUploadRequest) GetApiMethodName() string {
    return "alibaba.legal.suit.file.upload"
}

func (r AlibabaLegalSuitFileUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLegalSuitFileUploadRequest) SetFile(file []byte) error {
    r.file = file
    r.Set("file", file)
    return nil
}

func (r AlibabaLegalSuitFileUploadRequest) GetFile() []byte {
    return r.file
}

func (r *AlibabaLegalSuitFileUploadRequest) SetTimeStamp(timeStamp int64) error {
    r.timeStamp = timeStamp
    r.Set("time_stamp", timeStamp)
    return nil
}

func (r AlibabaLegalSuitFileUploadRequest) GetTimeStamp() int64 {
    return r.timeStamp
}

func (r *AlibabaLegalSuitFileUploadRequest) SetFileName(fileName string) error {
    r.fileName = fileName
    r.Set("file_name", fileName)
    return nil
}

func (r AlibabaLegalSuitFileUploadRequest) GetFileName() string {
    return r.fileName
}

func (r *AlibabaLegalSuitFileUploadRequest) SetFileSize(fileSize int64) error {
    r.fileSize = fileSize
    r.Set("file_size", fileSize)
    return nil
}

func (r AlibabaLegalSuitFileUploadRequest) GetFileSize() int64 {
    return r.fileSize
}

func (r *AlibabaLegalSuitFileUploadRequest) SetSignature(signature string) error {
    r.signature = signature
    r.Set("signature", signature)
    return nil
}

func (r AlibabaLegalSuitFileUploadRequest) GetSignature() string {
    return r.signature
}

