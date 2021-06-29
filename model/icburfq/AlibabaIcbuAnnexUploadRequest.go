package icburfq

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传附件获取附件files_str APIRequest
alibaba.icbu.annex.upload

上传附件获取附件files_str
*/
type AlibabaIcbuAnnexUploadRequest struct {
    model.Params

    // 文件名
    fileName   string 

    // 文件字节流
    fileInputStreamBytes   []*model.File 

    // 来源
    source   string 

}

func NewAlibabaIcbuAnnexUploadRequest() *AlibabaIcbuAnnexUploadRequest{
    return &AlibabaIcbuAnnexUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuAnnexUploadRequest) GetApiMethodName() string {
    return "alibaba.icbu.annex.upload"
}

func (r AlibabaIcbuAnnexUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIcbuAnnexUploadRequest) SetFileName(fileName string) error {
    r.fileName = fileName
    r.Set("file_name", fileName)
    return nil
}

func (r AlibabaIcbuAnnexUploadRequest) GetFileName() string {
    return r.fileName
}

func (r *AlibabaIcbuAnnexUploadRequest) SetFileInputStreamBytes(fileInputStreamBytes []*model.File) error {
    r.fileInputStreamBytes = fileInputStreamBytes
    r.Set("file_input_stream_bytes", fileInputStreamBytes)
    return nil
}

func (r AlibabaIcbuAnnexUploadRequest) GetFileInputStreamBytes() []*model.File {
    return r.fileInputStreamBytes
}

func (r *AlibabaIcbuAnnexUploadRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

func (r AlibabaIcbuAnnexUploadRequest) GetSource() string {
    return r.source
}

